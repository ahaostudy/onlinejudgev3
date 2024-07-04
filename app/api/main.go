package main

import (
	"github.com/ahaostudy/onlinejudge/app/api/mw/jwt"
	"github.com/ahaostudy/onlinejudge/app/api/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	rpc.InitClient()
	jwt.Init()

	RegisterRoute(h)
	h.Spin()
}
