package db

import (
	"context"
	ktdb "github.com/ahaostudy/kitextool/option/db"
	"github.com/ahaostudy/onlinejudge/app/user/model"
)

func GetById(ctx context.Context, id int64) (*model.User, error) {
	user := new(model.User)
	err := ktdb.DB().WithContext(ctx).Where("id = ?", id).First(user).Error
	return user, err
}

func GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user := new(model.User)
	err := ktdb.DB().WithContext(ctx).Where("username = ?", username).First(user).Error
	return user, err
}

func GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)
	err := ktdb.DB().WithContext(ctx).Where("email = ?", email).First(user).Error
	return user, err
}

func GetList(ctx context.Context, ids []int64) ([]*model.User, error) {
	var users []*model.User
	err := ktdb.DB().WithContext(ctx).Where("id in (?)", ids).Find(&users).Error
	return users, err
}

func Insert(ctx context.Context, user *model.User) error {
	return ktdb.DB().WithContext(ctx).Create(user).Error
}

func Update(ctx context.Context, id int64, user map[string]any) error {
	return ktdb.DB().WithContext(ctx).Model(new(model.User)).Where("id = ?", id).Updates(user).Error
}
