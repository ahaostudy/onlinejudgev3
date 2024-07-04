package main

import (
	"context"
	"github.com/ahaostudy/onlinejudge/kitex_gen/usersvc"
	"github.com/ahaostudy/onlinejudge/kitex_gen/usersvc/userservice"
	"os"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
)

func main() {
	cli := userservice.MustNewClient("user",
		client.WithHostPorts("127.0.0.1:8882"),
		client.WithTransportProtocol(transport.TTHeaderFramed),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)

	ctx := context.Background()

	result, err := cli.Login(ctx, &usersvc.LoginReq{
		Username: nil,
		Email:    nil,
		Password: nil,
		Captcha:  nil,
	})
	handleError(err)
	klog.Info(result)
}

func handleError(err error) {
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}
