package rpc

import (
	"sync"

	ktconf "github.com/ahaostudy/kitextool/conf"
	"github.com/ahaostudy/kitextool/suite/ktcsuite"
	usersvc "github.com/ahaostudy/onlinejudge/kitex_gen/usersvc/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
)

var (
	once sync.Once
	conf *ktconf.Default

	UserClient usersvc.Client
)

func InitClient() {
	once.Do(func() {
		UserClient = usersvc.MustNewClient("user",
			client.WithHostPorts("127.0.0.1:8882"),
			client.WithSuite(ktcsuite.NewKitexToolSuite(
				conf,
				ktcsuite.WithTransport(transport.TTHeaderFramed),
			)),
		)
	})
}
