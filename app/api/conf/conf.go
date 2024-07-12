package conf

import (
	"path/filepath"

	ktconf "github.com/ahaostudy/kitextool/conf"
	"github.com/kitex-contrib/config-nacos/nacos"
)

var conf = new(Conf)

func GetConf() *Conf {
	return conf
}

func init() {
	ktconf.LoadFiles(conf, filepath.Join("app", "api", "conf", "conf.yaml"))
	ktconf.ApplyDynamicConfig(ktconf.NewNacosConfigCenter(nacos.Options{}), &conf.ConfigCenter, conf.Server.Name, conf)
}

type Conf struct {
	ktconf.ClientConf
	Server       Server            `yaml:"server"`
	ConfigCenter ktconf.CenterConf `yaml:"config_center"`
}

type Server struct {
	Name string `yaml:"name"`
}
