package conf

import (
	"path/filepath"

	ktconf "github.com/ahaostudy/kitextool/conf"
	"github.com/kr/pretty"
)

var conf = new(Conf)

func GetConf() *Conf {
	return conf
}

func init() {
	ktconf.LoadFiles(conf,
		filepath.Join("conf", "conf.yaml"),
		filepath.Join("app", "user", "conf", "conf.yaml"),
	)
	_, _ = pretty.Printf("%+v\n", conf)
}

type Conf struct {
	ktconf.ServerConf
	Email Email `yaml:"email"`
	Auth  Auth  `yaml:"auth"`
}

type Email struct {
	SmtpAddr     string `yaml:"smtp_addr"`
	SmtpHost     string `yaml:"smtp_host"`
	EmailFrom    string `yaml:"email_from"`
	EmailAddress string `yaml:"email_address"`
	Auth         string `yaml:"auth"`
}

type Auth struct {
	Salt string `yaml:"salt"`
	Node int    `yaml:"node"`
}
