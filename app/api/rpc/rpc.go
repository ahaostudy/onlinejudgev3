package rpc

import (
	"sync"

	ktresolver "github.com/ahaostudy/kitextool/option/client/resolver"
	"github.com/ahaostudy/onlinejudge/app/api/conf"

	ktclient "github.com/ahaostudy/kitextool/suite/client"
	usersvc "github.com/ahaostudy/onlinejudge/kitex_gen/usersvc/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
)

var (
	once sync.Once

	UserCli usersvc.Client
)

func InitClient() {
	once.Do(func() {
		UserCli = usersvc.MustNewClient("user",
			client.WithSuite(ktclient.NewKitexToolSuite(
				conf.GetConf().ClientConf,
				ktclient.WithTransport(transport.TTHeaderFramed),
				ktresolver.WithResolver(ktresolver.NewNacosResolver),
			)),
		)
	})
}
