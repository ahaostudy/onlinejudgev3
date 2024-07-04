package handler

import (
	"context"

	"github.com/ahaostudy/onlinejudge/app/user/dal/db"
	"github.com/ahaostudy/onlinejudge/app/user/model"
	"github.com/ahaostudy/onlinejudge/app/user/pkg/sha256"
	"github.com/ahaostudy/onlinejudge/kitex_gen/usersvc"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func Login(ctx context.Context, req *usersvc.LoginReq) (resp *usersvc.LoginResp, err error) {
	if req.Username == nil && req.Email == nil {
		return nil, kerrors.NewBizStatusError(40021, "username and email is empty")
	}
	if req.Password == nil && req.Captcha == nil {
		return nil, kerrors.NewBizStatusError(40022, "identity verification failure")
	}

	var user *model.User
	if req.Email != nil {
		user, err = db.GetByEmail(ctx, *req.Email)
		if err != nil {
			return nil, kerrors.NewBizStatusError(50021, "failed to get user")
		}
		if req.Captcha != nil {
			if ok := CheckCaptcha(ctx, *req.Email, *req.Captcha); !ok {
				return nil, kerrors.NewBizStatusError(40022, "identity verification failure")
			}
		}
	} else {
		user, err = db.GetByUsername(ctx, *req.Username)
		if err != nil {
			return nil, kerrors.NewBizStatusError(50021, "failed to get user")
		}
	}
	if req.Password == nil || user.Password != sha256.Encrypt(*req.Password) {
		return nil, kerrors.NewBizStatusError(40022, "identity verification failure")
	}

	resp = &usersvc.LoginResp{UserId: user.Id}
	return
}
