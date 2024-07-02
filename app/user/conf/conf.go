package conf

import (
	ktconf "github.com/ahaostudy/kitextool/conf"
	"github.com/kr/pretty"
	"path/filepath"
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
	ktconf.Default
	Email Email `yaml:"email"`
	Auth  Auth  `yaml:"auth"`
}

type Email struct {
	Addr   string `yaml:"addr"`
	Host   string `yaml:"host"`
	From   string `yaml:"from"`
	Email  string `yaml:"email"`
	Auth   string `yaml:"auth"`
	Expire int    `yaml:"expire"`
}

type Auth struct {
	Salt string `yaml:"salt"`
	Node int    `yaml:"node"`
}
