package cache

import (
	"context"
	"encoding/json"
	"fmt"
	ktrdb "github.com/ahaostudy/kitextool/option/redis"
	"github.com/ahaostudy/onlinejudge/app/user/dal/db"
	"github.com/ahaostudy/onlinejudge/app/user/model"
	"time"
)

const (
	UserKey     = "user"
	IdKey       = "id"
	UsernameKey = "username"
	EmailKey    = "email"
)

type UserCache struct {
	ctx context.Context
}

func NewUserCache(ctx context.Context) *UserCache {
	return &UserCache{ctx: ctx}
}

func (c *UserCache) GetById(id int64) (*model.User, error) {
	key := fmt.Sprintf("%s:%s:%d", UserKey, IdKey, id)
	user, err := c.Get(key)
	if err != nil {
		user, err = db.GetById(c.ctx, id)
		if err != nil {
			return nil, err
		}
		_ = c.Set(key, user)
	}
	return user, nil
}

func (c *UserCache) GetByUsername(username string) (*model.User, error) {
	key := fmt.Sprintf("%s:%s:%s", UserKey, UsernameKey, username)
	user, err := c.Get(key)
	if err != nil {
		user, err = db.GetByUsername(c.ctx, username)
		if err != nil {
			return nil, err
		}
		_ = c.Set(key, user)
	}
	return user, nil
}

func (c *UserCache) GetByEmail(email string) (*model.User, error) {
	key := fmt.Sprintf("%s:%s:%s", UserKey, EmailKey, email)
	user, err := c.Get(key)
	if err != nil {
		user, err = db.GetByEmail(c.ctx, email)
		if err != nil {
			return nil, err
		}
		_ = c.Set(key, user)
	}
	return user, nil
}

func (c *UserCache) Set(key string, user *model.User) error {
	raw, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ktrdb.RDB().Set(c.ctx, key, raw, time.Hour).Err()
	return err
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
