package handler

import (
	"context"
	"errors"

	"github.com/ahaostudy/onlinejudge/app/user/dal/cache"

	"github.com/ahaostudy/onlinejudge/app/user/dal/db"
	"github.com/ahaostudy/onlinejudge/app/user/model"
	"github.com/ahaostudy/onlinejudge/app/user/pkg/sha256"
	"github.com/ahaostudy/onlinejudge/app/user/pkg/snowflake"
	"github.com/ahaostudy/onlinejudge/kitex_gen/usersvc"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/go-sql-driver/mysql"
)

func CreateUser(ctx context.Context, req *usersvc.CreateUserReq) (resp *usersvc.CreateUserResp, err error) {
	// verify that the parameters are legal
	// the username, email, and password fields cannot be empty
	// the user role be role_user or role_admin
	if len(req.Username) == 0 ||
		len(req.Email) == 0 ||
		len(req.Password) == 0 ||
		(req.Role != int32(model.RoleUser) && req.Role != int32(model.RoleAdmin)) {
		return nil, kerrors.NewBizStatusError(40031, "invalid param")
	}
	loggedInUser, err := cache.NewUserCache(ctx).GetById(req.LoggedInId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50031, "server failure")
	}
	if loggedInUser.Role != model.RoleAdmin {
		return nil, kerrors.NewBizStatusError(40032, "forbidden")
	}
	id := snowflake.Generate().Int64()
	user := &model.User{
		Id:        id,
		Nickname:  req.Nickname,
		Username:  req.Username,
		Password:  sha256.Encrypt(req.Password),
		Email:     req.Email,
		Avatar:    req.Avatar,
		Signature: req.Signature,
		Role:      model.UserRole(req.Role),
	}
	if user.Signature == "" {
		user.Signature = model.DefaultSignature(req.Username)
	}
	err = db.Insert(ctx, user)
	if err != nil {
		var e *mysql.MySQLError
		if errors.As(err, &e) && e.Number == 1062 {
			return nil, kerrors.NewBizStatusError(40033, "the user already exists")
		}
		return nil, kerrors.NewBizStatusError(50032, "failed to insert user")
	}
	resp = &usersvc.CreateUserResp{UserId: id}
	return
}
