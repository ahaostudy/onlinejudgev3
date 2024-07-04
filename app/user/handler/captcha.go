package handler

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ahaostudy/onlinejudge/app/user/dal/cache"
	"github.com/ahaostudy/onlinejudge/app/user/pkg/email"
	"github.com/ahaostudy/onlinejudge/kitex_gen/usersvc"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func GenCaptcha(ctx context.Context, req *usersvc.GenCaptchaReq) (resp *usersvc.GenCaptchaResp, err error) {
	// generate random captcha
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	captcha := fmt.Sprintf("%d", r.Intn(900000)+100000)

	errCh := make(chan error)
	go func() {
		errCh <- email.SendCaptcha(captcha, req.Email)
	}()
	go func() {
		errCh <- cache.NewCaptchaCache(ctx).Set(req.Email, captcha)
	}()

	if <-errCh == nil && <-errCh == nil {
		return &usersvc.GenCaptchaResp{Captcha: captcha}, nil
	} else {
		return nil, kerrors.NewBizStatusError(50071, "failed to generate captcha")
	}
}
