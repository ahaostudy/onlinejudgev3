package user

import (
	"context"
	"net/http"

	"github.com/ahaostudy/onlinejudge/app/api/dto"
	"github.com/ahaostudy/onlinejudge/app/api/mw/jwt"
	"github.com/ahaostudy/onlinejudge/app/api/rpc"
	"github.com/ahaostudy/onlinejudge/kitex_gen/usersvc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dto.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, dto.RegisterResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}

	result, err := rpc.UserCli.Register(ctx, &usersvc.RegisterReq{
		Email:    req.Email,
		Captcha:  req.Captcha,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusOK, dto.RegisterResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}

	// TODO: due to the expiration of the captcha, it will result in an additional invalid validation process when logging in
	jwt.Middleware.LoginHandler(ctx, c)
	token := c.GetString("token")
	expire := c.GetString("expire")

	c.JSON(http.StatusOK, dto.RegisterResp{BaseResp: dto.SuccessResp(), UserId: result.UserId, Token: token, Expire: expire})
}

func Login(ctx context.Context, c *app.RequestContext) {
	token := c.GetString("token")
	userId := c.GetInt64(jwt.IdentityKey)
	expire := c.GetString("expire")
	c.JSON(http.StatusOK, dto.LoginResp{BaseResp: dto.SuccessResp(), UserId: userId, Token: token, Expire: expire})
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
	result, err := rpc.UserCli.GetUser(ctx, req)
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
	result, err := rpc.UserCli.CreateUser(ctx, &usersvc.CreateUserReq{
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
	_, err = rpc.UserCli.UpdateUser(ctx, &usersvc.UpdateUserReq{
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

	result, err := rpc.UserCli.GenCaptcha(ctx, &usersvc.GenCaptchaReq{Email: req.Email})
	if err != nil {
		c.JSON(consts.StatusOK, dto.GetCaptchaResp{BaseResp: dto.BuildBaseResp(err)})
		return
	}
	hlog.CtxInfof(ctx, "captcha: %s", result.Captcha)
	c.JSON(http.StatusOK, dto.GetCaptchaResp{BaseResp: dto.SuccessResp()})
}
