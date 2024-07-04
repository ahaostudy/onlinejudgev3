package sha256

import (
	"testing"

	"github.com/ahaostudy/onlinejudge/app/user/conf"
)

func TestEncrypt(t *testing.T) {
	conf.GetConf().Auth.Salt = "DF8%sd2%Df2^3fIN98"
	pwd := "123456"
	encrypt := Encrypt(pwd)
	t.Log(encrypt)
}
