package main

import (
	"context"

	"github.com/ahaostudy/onlinejudge/app/user/handler"
	base "github.com/ahaostudy/onlinejudge/kitex_gen/base"
	usersvc "github.com/ahaostudy/onlinejudge/kitex_gen/usersvc"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *usersvc.RegisterReq) (resp *usersvc.RegisterResp, err error) {
	return handler.Register(ctx, req)
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *usersvc.LoginReq) (resp *usersvc.LoginResp, err error) {
	return handler.Login(ctx, req)
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *usersvc.CreateUserReq) (resp *usersvc.CreateUserResp, err error) {
	return handler.CreateUser(ctx, req)
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *usersvc.UpdateUserReq) (resp *base.Empty, err error) {
	return handler.UpdateUser(ctx, req)
}

// GenCaptcha implements the UserServiceImpl interface.
func (s *UserServiceImpl) GenCaptcha(ctx context.Context, req *usersvc.GenCaptchaReq) (resp *usersvc.GenCaptchaResp, err error) {
	return handler.GenCaptcha(ctx, req)
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *usersvc.GetUserReq) (resp *usersvc.GetUserResp, err error) {
	return handler.GetUser(ctx, req)
}

// GetUserList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserList(ctx context.Context, req *usersvc.GetUserListReq) (resp *usersvc.GetUserListResp, err error) {
	return handler.GetUserList(ctx, req)
}

// UploadAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) UploadAvatar(ctx context.Context, req *usersvc.UploadAvatarReq) (resp *usersvc.UploadAvatarResp, err error) {
	return handler.UploadAvatar(ctx, req)
}

// DownloadAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) DownloadAvatar(ctx context.Context, req *usersvc.DownloadAvatarReq) (resp *usersvc.DownloadAvatarResp, err error) {
	return handler.DownloadAvatar(ctx, req)
}

// DeleteAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteAvatar(ctx context.Context, req *usersvc.DeleteAvatarReq) (resp *base.Empty, err error) {
	return handler.DeleteAvatar(ctx, req)
}
