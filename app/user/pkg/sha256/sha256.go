package sha256

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/ahaostudy/onlinejudge/app/user/conf"
)

func Encrypt(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password + conf.GetConf().Auth.Salt))
	return hex.EncodeToString(hash.Sum(nil))
}
