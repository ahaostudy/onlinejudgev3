package handler

import (
	"context"
	"errors"

	"github.com/ahaostudy/onlinejudge/app/user/dal/cache"
	"github.com/ahaostudy/onlinejudge/app/user/dal/db"
	"github.com/ahaostudy/onlinejudge/app/user/model"
	"github.com/ahaostudy/onlinejudge/app/user/pkg/sha256"
	"github.com/ahaostudy/onlinejudge/kitex_gen/base"
	"github.com/ahaostudy/onlinejudge/kitex_gen/usersvc"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/go-sql-driver/mysql"
)

func UpdateUser(ctx context.Context, req *usersvc.UpdateUserReq) (resp *base.Empty, err error) {
	userCache := cache.NewUserCache(ctx)
	loggedInUser, err := userCache.GetById(req.LoggedInId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50041, "failed to get user")
	}
	// ordinary user cannot update other user information and their identity information
	if loggedInUser.Role == model.RoleUser && (req.LoggedInId != req.Id || req.Role != nil) {
		return nil, kerrors.NewBizStatusError(40041, "forbidden")
	}

	updateMap := make(map[string]interface{})
	if req.Username != nil {
		updateMap["username"] = *req.Username
	}
	if req.Email != nil {
		updateMap["email"] = *req.Email
	}
	if req.Nickname != nil {
		updateMap["nickname"] = *req.Nickname
	}
	if req.Password != nil {
		updateMap["password"] = sha256.Encrypt(*req.Password)
	}
	if req.Avatar != nil {
		updateMap["avatar"] = *req.Avatar
	}
	if req.Signature != nil {
		updateMap["signature"] = *req.Signature
	}
	if req.Role != nil {
		updateMap["role"] = model.UserRole(*req.Role)
	}

	user, err := db.Update(ctx, req.Id, updateMap)
	if err != nil {
		var me *mysql.MySQLError
		if errors.As(err, &me) && me.Number == 1062 {
			return nil, kerrors.NewBizStatusError(40042, "the user already exists")
		}
		return nil, kerrors.NewBizStatusError(50042, "failed to update user")
	}

	_ = userCache.SetById(req.Id, user)
	return
}
