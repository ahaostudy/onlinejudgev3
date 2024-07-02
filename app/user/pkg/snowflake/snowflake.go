package snowflake

import (
	"github.com/ahaostudy/onlinejudge/app/user/conf"
	"github.com/bwmarrin/snowflake"
	"sync"
)

var (
	node *snowflake.Node
	once sync.Once
)

func initSnowFlake() {
	n, err := snowflake.NewNode(int64(conf.GetConf().Auth.Node))
	if err != nil {
		panic(err)
	}
	node = n
}

func Generate() snowflake.ID {
	once.Do(initSnowFlake)
	return node.Generate()
}
