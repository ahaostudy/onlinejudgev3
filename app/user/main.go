package main

import (
	"log"

	"github.com/ahaostudy/onlinejudge/app/user/dal/db"

	ktconf "github.com/ahaostudy/kitextool/conf"
	ktdb "github.com/ahaostudy/kitextool/option/server/db"
	ktrdb "github.com/ahaostudy/kitextool/option/server/redis"
	ktregistry "github.com/ahaostudy/kitextool/option/server/registry"
	ktserver "github.com/ahaostudy/kitextool/suite/server"
	"github.com/ahaostudy/onlinejudge/app/user/conf"
	usersvc "github.com/ahaostudy/onlinejudge/kitex_gen/usersvc/userservice"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/config-nacos/nacos"
)

func main() {
	svr := usersvc.NewServer(new(UserServiceImpl),
		server.WithSuite(ktserver.NewKitexToolSuite(
			conf.GetConf(),
			ktserver.WithTransport(transport.TTHeaderFramed),
			ktserver.WithDynamicConfig(ktconf.NewNacosConfigCenter(nacos.Options{})),
			ktregistry.WithRegistry(ktregistry.NewNacosRegistry()),
			ktdb.WithDB(ktdb.NewMySQLDial()),
			ktrdb.WithRedis(),
		)),
	)

	db.AutoMigrate()

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
