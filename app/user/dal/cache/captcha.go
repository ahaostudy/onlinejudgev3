package cache

import (
	"context"
	"fmt"
	"time"

	ktrdb "github.com/ahaostudy/kitextool/option/redis"
)

const (
	CaptchaKey = "captcha"
)

type CaptchaCache struct {
	ctx context.Context
}

func (c *CaptchaCache) Set(email, captcha string) error {
	key := fmt.Sprintf("%s:%s", CaptchaKey, email)
	return ktrdb.RDB().Set(c.ctx, key, captcha, 5*time.Minute).Err()
}

func (c *CaptchaCache) Get(email string) (string, error) {
	key := fmt.Sprintf("%s:%s", CaptchaKey, email)
	captcha, err := ktrdb.RDB().Get(c.ctx, key).Result()
	if err != nil {
		return "", err
	}
	return captcha, nil
}

func (c *CaptchaCache) Del(email string) error {
	key := fmt.Sprintf("%s:%s", CaptchaKey, email)
	return ktrdb.RDB().Del(c.ctx, key).Err()
}

func NewCaptchaCache(ctx context.Context) *CaptchaCache {
	return &CaptchaCache{ctx: ctx}
}
