package main

import (
	"log"

	ktconf "github.com/ahaostudy/kitextool/conf"
	ktrdb "github.com/ahaostudy/kitextool/option/server/redis"
	ktregistry "github.com/ahaostudy/kitextool/option/server/registry"
	ktserver "github.com/ahaostudy/kitextool/suite/server"
	"github.com/ahaostudy/onlinejudge/app/judge/conf"
	judgesvc "github.com/ahaostudy/onlinejudge/kitex_gen/judgesvc/judgeservice"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/config-nacos/nacos"
	nacosserver "github.com/kitex-contrib/config-nacos/server"
)

func main() {
	nacosCenter := ktconf.NewNacosConfigCenter(nacos.Options{})
	nacosCenter.Init(&conf.GetConf().ConfigCenter)

	svr := judgesvc.NewServer(new(JudgeServiceImpl),
		server.WithSuite(ktserver.NewKitexToolSuite(
			conf.GetConf(),
			ktserver.WithTransport(transport.TTHeaderFramed),
			ktserver.WithDynamicConfig(nacosCenter),
			ktregistry.WithRegistry(ktregistry.NewNacosRegistry()),
			ktrdb.WithRedis(),
		)),
		server.WithSuite(nacosserver.NewSuite(conf.GetConf().Server.Name, nacosCenter.Client())),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
