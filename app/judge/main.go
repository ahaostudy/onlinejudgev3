package main

import (
	ktconf "github.com/ahaostudy/kitextool/conf"
	ktrdb "github.com/ahaostudy/kitextool/option/redis"
	"github.com/ahaostudy/kitextool/suite/ktssuite"
	"github.com/ahaostudy/onlinejudge/app/judge/conf"
	judgesvc "github.com/ahaostudy/onlinejudge/kitex_gen/judgesvc/judgeservice"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/config-nacos/nacos"
	nacosserver "github.com/kitex-contrib/config-nacos/server"
	"log"
)

func main() {
	nacosCenter := ktconf.NewNacosConfigCenter(nacos.Options{})
	nacosCenter.InitClient(&conf.GetConf().ConfigCenter)

	svr := judgesvc.NewServer(new(JudgeServiceImpl),
		server.WithSuite(ktssuite.NewKitexToolSuite(
			conf.GetConf(),
			ktssuite.WithDynamicConfig(nacosCenter),
			// ktregistry.WithRegistry(ktregistry.NewNacosRegistry()),
			ktrdb.WithRedis(),
		)),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithSuite(nacosserver.NewSuite(conf.GetConf().Server.Name, nacosCenter.Client())),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
