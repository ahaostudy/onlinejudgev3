// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	base "github.com/ahaostudy/onlinejudge/kitex_gen/base"
	usersvc "github.com/ahaostudy/onlinejudge/kitex_gen/usersvc"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newUserServiceRegisterArgs,
		newUserServiceRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newUserServiceLoginArgs,
		newUserServiceLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateUser": kitex.NewMethodInfo(
		createUserHandler,
		newUserServiceCreateUserArgs,
		newUserServiceCreateUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Update": kitex.NewMethodInfo(
		updateHandler,
		newUserServiceUpdateArgs,
		newUserServiceUpdateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GenCaptcha": kitex.NewMethodInfo(
		genCaptchaHandler,
		newUserServiceGenCaptchaArgs,
		newUserServiceGenCaptchaResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetPermission": kitex.NewMethodInfo(
		getPermissionHandler,
		newUserServiceGetPermissionArgs,
		newUserServiceGetPermissionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetUser": kitex.NewMethodInfo(
		getUserHandler,
		newUserServiceGetUserArgs,
		newUserServiceGetUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetUserList": kitex.NewMethodInfo(
		getUserListHandler,
		newUserServiceGetUserListArgs,
		newUserServiceGetUserListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UploadAvatar": kitex.NewMethodInfo(
		uploadAvatarHandler,
		newUserServiceUploadAvatarArgs,
		newUserServiceUploadAvatarResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DownloadAvatar": kitex.NewMethodInfo(
		downloadAvatarHandler,
		newUserServiceDownloadAvatarArgs,
		newUserServiceDownloadAvatarResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteAvatar": kitex.NewMethodInfo(
		deleteAvatarHandler,
		newUserServiceDeleteAvatarArgs,
		newUserServiceDeleteAvatarResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*usersvc.UserService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "usersvc",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usersvc.UserServiceRegisterArgs)
	realResult := result.(*usersvc.UserServiceRegisterResult)
	success, err := handler.(usersvc.UserService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return usersvc.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return usersvc.NewUserServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usersvc.UserServiceLoginArgs)
	realResult := result.(*usersvc.UserServiceLoginResult)
	success, err := handler.(usersvc.UserService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return usersvc.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return usersvc.NewUserServiceLoginResult()
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usersvc.UserServiceCreateUserArgs)
	realResult := result.(*usersvc.UserServiceCreateUserResult)
	success, err := handler.(usersvc.UserService).CreateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCreateUserArgs() interface{} {
	return usersvc.NewUserServiceCreateUserArgs()
}

func newUserServiceCreateUserResult() interface{} {
	return usersvc.NewUserServiceCreateUserResult()
}

func updateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usersvc.UserServiceUpdateArgs)
	realResult := result.(*usersvc.UserServiceUpdateResult)
	success, err := handler.(usersvc.UserService).Update(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateArgs() interface{} {
	return usersvc.NewUserServiceUpdateArgs()
}

func newUserServiceUpdateResult() interface{} {
	return usersvc.NewUserServiceUpdateResult()
}

func genCaptchaHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usersvc.UserServiceGenCaptchaArgs)
	realResult := result.(*usersvc.UserServiceGenCaptchaResult)
	success, err := handler.(usersvc.UserService).GenCaptcha(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGenCaptchaArgs() interface{} {
	return usersvc.NewUserServiceGenCaptchaArgs()
}

func newUserServiceGenCaptchaResult() interface{} {
	return usersvc.NewUserServiceGenCaptchaResult()
}

func getPermissionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usersvc.UserServiceGetPermissionArgs)
	realResult := result.(*usersvc.UserServiceGetPermissionResult)
	success, err := handler.(usersvc.UserService).GetPermission(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetPermissionArgs() interface{} {
	return usersvc.NewUserServiceGetPermissionArgs()
}

func newUserServiceGetPermissionResult() interface{} {
	return usersvc.NewUserServiceGetPermissionResult()
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usersvc.UserServiceGetUserArgs)
	realResult := result.(*usersvc.UserServiceGetUserResult)
	success, err := handler.(usersvc.UserService).GetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserArgs() interface{} {
	return usersvc.NewUserServiceGetUserArgs()
}

func newUserServiceGetUserResult() interface{} {
	return usersvc.NewUserServiceGetUserResult()
}

func getUserListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usersvc.UserServiceGetUserListArgs)
	realResult := result.(*usersvc.UserServiceGetUserListResult)
	success, err := handler.(usersvc.UserService).GetUserList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserListArgs() interface{} {
	return usersvc.NewUserServiceGetUserListArgs()
}

func newUserServiceGetUserListResult() interface{} {
	return usersvc.NewUserServiceGetUserListResult()
}

func uploadAvatarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usersvc.UserServiceUploadAvatarArgs)
	realResult := result.(*usersvc.UserServiceUploadAvatarResult)
	success, err := handler.(usersvc.UserService).UploadAvatar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUploadAvatarArgs() interface{} {
	return usersvc.NewUserServiceUploadAvatarArgs()
}

func newUserServiceUploadAvatarResult() interface{} {
	return usersvc.NewUserServiceUploadAvatarResult()
}

func downloadAvatarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usersvc.UserServiceDownloadAvatarArgs)
	realResult := result.(*usersvc.UserServiceDownloadAvatarResult)
	success, err := handler.(usersvc.UserService).DownloadAvatar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceDownloadAvatarArgs() interface{} {
	return usersvc.NewUserServiceDownloadAvatarArgs()
}

func newUserServiceDownloadAvatarResult() interface{} {
	return usersvc.NewUserServiceDownloadAvatarResult()
}

func deleteAvatarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usersvc.UserServiceDeleteAvatarArgs)
	realResult := result.(*usersvc.UserServiceDeleteAvatarResult)
	success, err := handler.(usersvc.UserService).DeleteAvatar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceDeleteAvatarArgs() interface{} {
	return usersvc.NewUserServiceDeleteAvatarArgs()
}

func newUserServiceDeleteAvatarResult() interface{} {
	return usersvc.NewUserServiceDeleteAvatarResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *usersvc.RegisterReq) (r *usersvc.RegisterResp, err error) {
	var _args usersvc.UserServiceRegisterArgs
	_args.Req = req
	var _result usersvc.UserServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *usersvc.LoginReq) (r *usersvc.LoginResp, err error) {
	var _args usersvc.UserServiceLoginArgs
	_args.Req = req
	var _result usersvc.UserServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateUser(ctx context.Context, req *usersvc.CreateUserReq) (r *usersvc.CreateUserResp, err error) {
	var _args usersvc.UserServiceCreateUserArgs
	_args.Req = req
	var _result usersvc.UserServiceCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Update(ctx context.Context, req *usersvc.UpdateUserReq) (r *base.Empty, err error) {
	var _args usersvc.UserServiceUpdateArgs
	_args.Req = req
	var _result usersvc.UserServiceUpdateResult
	if err = p.c.Call(ctx, "Update", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GenCaptcha(ctx context.Context, req *usersvc.GenCaptchaReq) (r *usersvc.GenCaptchaResp, err error) {
	var _args usersvc.UserServiceGenCaptchaArgs
	_args.Req = req
	var _result usersvc.UserServiceGenCaptchaResult
	if err = p.c.Call(ctx, "GenCaptcha", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPermission(ctx context.Context, req *usersvc.GetPermissionReq) (r *usersvc.GetPermissionResp, err error) {
	var _args usersvc.UserServiceGetPermissionArgs
	_args.Req = req
	var _result usersvc.UserServiceGetPermissionResult
	if err = p.c.Call(ctx, "GetPermission", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUser(ctx context.Context, req *usersvc.GetUserReq) (r *usersvc.GetUserResp, err error) {
	var _args usersvc.UserServiceGetUserArgs
	_args.Req = req
	var _result usersvc.UserServiceGetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserList(ctx context.Context, req *usersvc.GetUserListReq) (r *usersvc.GetUserListResp, err error) {
	var _args usersvc.UserServiceGetUserListArgs
	_args.Req = req
	var _result usersvc.UserServiceGetUserListResult
	if err = p.c.Call(ctx, "GetUserList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UploadAvatar(ctx context.Context, req *usersvc.UploadAvatarReq) (r *usersvc.UploadAvatarResp, err error) {
	var _args usersvc.UserServiceUploadAvatarArgs
	_args.Req = req
	var _result usersvc.UserServiceUploadAvatarResult
	if err = p.c.Call(ctx, "UploadAvatar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DownloadAvatar(ctx context.Context, req *usersvc.DownloadAvatarReq) (r *usersvc.DownloadAvatarResp, err error) {
	var _args usersvc.UserServiceDownloadAvatarArgs
	_args.Req = req
	var _result usersvc.UserServiceDownloadAvatarResult
	if err = p.c.Call(ctx, "DownloadAvatar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteAvatar(ctx context.Context, req *usersvc.DeleteAvatarReq) (r *base.Empty, err error) {
	var _args usersvc.UserServiceDeleteAvatarArgs
	_args.Req = req
	var _result usersvc.UserServiceDeleteAvatarResult
	if err = p.c.Call(ctx, "DeleteAvatar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
