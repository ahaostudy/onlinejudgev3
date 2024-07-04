package handler

import (
	"context"
	"github.com/ahaostudy/onlinejudge/app/user/dal/cache"
	"github.com/ahaostudy/onlinejudge/app/user/dal/db"
	"github.com/ahaostudy/onlinejudge/app/user/model"
	"github.com/ahaostudy/onlinejudge/kitex_gen/usersvc"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func GetUser(ctx context.Context, req *usersvc.GetUserReq) (resp *usersvc.GetUserResp, err error) {
	var user *model.User
	if req.Id != nil {
		userCache := cache.NewUserCache(ctx)
		user, err = userCache.GetById(*req.Id)
		if err != nil {
			return nil, kerrors.NewBizStatusError(40061, "user not exists")
		}
	} else if req.Username != nil {
		user, err = db.GetByUsername(ctx, *req.Username)
		if err != nil {
			return nil, kerrors.NewBizStatusError(40061, "user not exists")
		}
	} else {
		return nil, kerrors.NewBizStatusError(40062, "invalid param")
	}
	resp = &usersvc.GetUserResp{
		User: &usersvc.User{
			Id:        user.Id,
			Nickname:  user.Nickname,
			Username:  user.Username,
			Email:     user.Email,
			Avatar:    user.Avatar,
			Signature: user.Signature,
			Role:      int32(user.Role),
		},
	}
	return
}
