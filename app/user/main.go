package main

import (
	usersvc "github.com/ahaostudy/onlinejudge/kitex_gen/usersvc/userservice"
	"log"
)

func main() {
	svr := usersvc.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
