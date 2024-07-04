package user

import (
	"context"
	"github.com/ahaostudy/onlinejudge/app/api/dto"
	"github.com/ahaostudy/onlinejudge/app/api/mw/jwt"
	"github.com/ahaostudy/onlinejudge/app/api/rpc"
	"github.com/ahaostudy/onlinejudge/kitex_gen/usersvc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dto.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, dto.RegisterResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}

	result, err := rpc.UserClient.Register(ctx, &usersvc.RegisterReq{
		Email:    req.Email,
		Captcha:  req.Captcha,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusOK, dto.RegisterResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}

	c.JSON(http.StatusOK, dto.RegisterResp{BaseResp: dto.SuccessResp(), UserId: result.UserId})
}

func GetUser(ctx context.Context, c *app.RequestContext) {
	userId := c.GetInt64(jwt.IdentityKey)
	getUser(ctx, c, &usersvc.GetUserReq{Id: &userId})
}

func GetUserById(ctx context.Context, c *app.RequestContext) {
	var req dto.GetUserByIdReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, dto.GetUserResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}
	getUser(ctx, c, &usersvc.GetUserReq{Id: &req.Id})
}

func GetUserByUsername(ctx context.Context, c *app.RequestContext) {
	var req dto.GetUserByUsernameReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, dto.GetUserResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}
	getUser(ctx, c, &usersvc.GetUserReq{Username: &req.Username})
}

func getUser(ctx context.Context, c *app.RequestContext, req *usersvc.GetUserReq) {
	result, err := rpc.UserClient.GetUser(ctx, req)
	if err != nil {
		c.JSON(consts.StatusOK, dto.GetUserResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}
	user := &dto.User{
		Id:        result.User.Id,
		Nickname:  result.User.Nickname,
		Username:  result.User.Username,
		Email:     result.User.Email,
		Avatar:    result.User.Avatar,
		Signature: result.User.Signature,
		Role:      result.User.Role,
	}
	c.JSON(consts.StatusOK, dto.GetUserResp{BaseResp: dto.SuccessResp(), User: user})
}

func CreateUser(ctx context.Context, c *app.RequestContext) {
	var req dto.CreateUserReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, dto.CreateUserResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}
	result, err := rpc.UserClient.CreateUser(ctx, &usersvc.CreateUserReq{
		Nickname:  req.Nickname,
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		Avatar:    req.Avatar,
		Signature: req.Signature,
		Role:      req.Role,
	})
	if err != nil {
		c.JSON(consts.StatusOK, dto.CreateUserResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}
	c.JSON(http.StatusOK, dto.CreateUserResp{BaseResp: dto.SuccessResp(), UserId: result.UserId})
}

func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var req dto.UpdateUserReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, dto.UpdateUserResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}
	loggedInId := c.GetInt64(jwt.IdentityKey)
	_, err = rpc.UserClient.UpdateUser(ctx, &usersvc.UpdateUserReq{
		Id:         req.Id,
		LoggedInId: loggedInId,
		Nickname:   req.Nickname,
		Username:   req.Username,
		Password:   req.Password,
		Email:      req.Email,
		Avatar:     req.Avatar,
		Signature:  req.Signature,
		Role:       req.Role,
	})
	if err != nil {
		c.JSON(consts.StatusOK, dto.UpdateUserResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}
	c.JSON(consts.StatusOK, dto.UpdateUserResp{BaseResp: dto.SuccessResp()})
}

func GetCaptcha(ctx context.Context, c *app.RequestContext) {
	var req dto.GetCaptchaReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, dto.LoginResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}

	result, err := rpc.UserClient.GenCaptcha(ctx, &usersvc.GenCaptchaReq{Email: req.Email})
	if err != nil {
		c.JSON(consts.StatusOK, dto.GetCaptchaResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}
	hlog.CtxInfof(ctx, "captcha: %s", result.Captcha)
	c.JSON(http.StatusOK, dto.GetCaptchaResp{BaseResp: dto.SuccessResp()})
}
