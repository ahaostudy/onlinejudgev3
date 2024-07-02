package handler

import (
	"context"
	"github.com/ahaostudy/onlinejudge/app/user/dal/cache"
	"github.com/ahaostudy/onlinejudge/app/user/dal/db"
	"github.com/ahaostudy/onlinejudge/app/user/model"
	"github.com/ahaostudy/onlinejudge/app/user/pkg/email"
	"github.com/ahaostudy/onlinejudge/app/user/pkg/sha256"
	"github.com/ahaostudy/onlinejudge/app/user/pkg/snowflake"
	"github.com/ahaostudy/onlinejudge/kitex_gen/usersvc"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"regexp"
)

func Register(ctx context.Context, req *usersvc.RegisterReq) (resp *usersvc.RegisterResp, err error) {
	userCache := cache.NewUserCache(ctx)

	if !CheckPassword(req.GetPassword()) {
		return nil, kerrors.NewBizStatusError(40011, "illegal password")
	}
	if !CheckCaptcha(ctx, req.GetEmail(), req.GetCaptcha()) {
		return nil, kerrors.NewBizStatusError(40012, "invalid captcha")
	}
	_, err = userCache.GetByEmail(req.Email)
	if err != nil {
		return nil, kerrors.NewBizStatusError(40013, "email exist")
	}

	// extract username
	username, ok := email.ExtractUsernameFromEmail(req.Email)
	if !ok {
		return nil, kerrors.NewBizStatusError(40014, "invalid email")
	}

	id := snowflake.Generate().Int64()
	user := model.User{
		ID:        id,
		Email:     req.Email,
		Nickname:  username,
		Username:  username,
		Password:  sha256.Encrypt(req.Password),
		Signature: model.DefaultSignature(username),
		Role:      model.RoleUser,
	}
	err = db.Insert(ctx, &user)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50011, "failed to insert user")
	}
	resp = &usersvc.RegisterResp{UserId: id}
	return
}

// CheckPassword check password security
func CheckPassword(password string) bool {
	var (
		hasMinLength   = len(password) >= 6
		hasUpperCase   = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLowerCase   = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber      = regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpecialChar = regexp.MustCompile(`[\~\!\?\@\#\$\%\^\&\*\_\-\+\=\(\)\[\]\{\}\>\<\/\\\|\"\'\.\,\:\;]`).MatchString(password)
	)
	return hasMinLength && hasUpperCase && hasLowerCase && (hasNumber || hasSpecialChar)
}

// CheckCaptcha check captcha correctness
func CheckCaptcha(ctx context.Context, email, captcha string) bool {
	captchaCache := cache.NewCaptchaCache(ctx)
	cpt, err := captchaCache.Get(email)
	if err != nil {
		return false
	}
	_ = captchaCache.Del(email)
	if cpt != captcha {
		return false
	}
	return true
}
