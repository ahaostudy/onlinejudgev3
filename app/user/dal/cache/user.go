package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	ktrdb "github.com/ahaostudy/kitextool/option/redis"
	"github.com/ahaostudy/onlinejudge/app/user/dal/db"
	"github.com/ahaostudy/onlinejudge/app/user/model"
	"gorm.io/gorm"
)

const (
	UserKey = "user"
	IdKey   = "id"
)

var (
	NotExistsUserId int64 = -1
	NotExistsUser         = &model.User{Id: NotExistsUserId}
)

type UserCache struct {
	ctx context.Context
}

func (c *UserCache) GetById(id int64) (*model.User, error) {
	key := fmt.Sprintf("%s:%s:%d", UserKey, IdKey, id)
	user, err := c.Get(key)
	if err != nil {
		user, err = db.GetById(c.ctx, id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				_ = c.Set(key, NotExistsUser)
			}
			return nil, err
		}
		_ = c.Set(key, user)
		return user, nil
	}
	if user.Id == NotExistsUserId {
		return nil, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (c *UserCache) SetById(id int64, user *model.User) error {
	key := fmt.Sprintf("%s:%s:%d", UserKey, IdKey, id)
	return c.Set(key, user)
}

func (c *UserCache) DelById(id int64) error {
	key := fmt.Sprintf("%s:%s:%d", UserKey, IdKey, id)
	return c.Del(key)
}

func (c *UserCache) Get(key string) (*model.User, error) {
	result := ktrdb.RDB().Get(c.ctx, key)
	err := result.Err()
	if err != nil {
		return nil, err
	}
	raw, err := result.Bytes()
	if err != nil {
		return nil, err
	}
	user := new(model.User)
	err = json.Unmarshal(raw, &user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (c *UserCache) Set(key string, user *model.User) error {
	raw, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ktrdb.RDB().Set(c.ctx, key, raw, time.Hour).Err()
	return err
}

func (c *UserCache) Del(key string) error {
	return ktrdb.RDB().Del(c.ctx, key).Err()
}

func NewUserCache(ctx context.Context) *UserCache {
	return &UserCache{ctx: ctx}
}
