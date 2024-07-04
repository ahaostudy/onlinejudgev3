package handler

import (
	"context"
	"github.com/ahaostudy/onlinejudge/app/user/dal/db"

	"github.com/ahaostudy/onlinejudge/kitex_gen/usersvc"
)

func GetUserList(ctx context.Context, req *usersvc.GetUserListReq) (resp *usersvc.GetUserListResp, err error) {
	users, err := db.GetList(ctx, req.UserIdList)
	var usersResp []*usersvc.User
	for _, user := range users {
		usersResp = append(usersResp, &usersvc.User{
			Id:        user.Id,
			Nickname:  user.Nickname,
			Username:  user.Username,
			Email:     user.Email,
			Avatar:    user.Avatar,
			Signature: user.Signature,
			Role:      int32(user.Role),
		})
	}
	resp = &usersvc.GetUserListResp{UserList: usersResp}
	return
}
