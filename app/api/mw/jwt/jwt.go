package jwt

import (
	"context"
	"errors"
	"time"

	"github.com/ahaostudy/onlinejudge/app/api/dto"
	"github.com/ahaostudy/onlinejudge/app/api/rpc"
	"github.com/ahaostudy/onlinejudge/kitex_gen/usersvc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hertz-contrib/jwt"
)

var (
	Middleware  *jwt.HertzJWTMiddleware
	IdentityKey = "user_id"
)

func Init() {
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			if userFloatId, ok := claims[IdentityKey].(float64); ok {
				return int64(userFloatId)
			}
			return 0
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var req dto.LoginReq
			err := c.BindAndValidate(&req)
			if err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			result, err := rpc.UserCli.Login(ctx, &usersvc.LoginReq{
				Username: req.Username,
				Email:    req.Email,
				Password: req.Password,
				Captcha:  req.Captcha,
			})
			if err != nil {
				if bizErr, ok := kerrors.FromBizStatusError(err); ok {
					return nil, errors.New(bizErr.BizMessage())
				}
				return nil, errors.New("server failure")
			}
			return result.UserId, nil
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			var statusCode int32
			if code != consts.StatusOK {
				statusCode = int32(code)
			}
			c.JSON(code, dto.BaseResp{StatusCode: statusCode, StatusMsg: message})
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.Set("token", token)
			c.Set("expire", expire.Format(time.RFC3339))
		},
		RefreshResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(code, dto.LoginResp{BaseResp: dto.SuccessResp(), Token: token, Expire: expire.Format(time.RFC3339)})
		},
	})
	if err != nil {
		panic("[JWT] error:" + err.Error())
	}
	Middleware = authMiddleware
}
