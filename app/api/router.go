package main

import (
	"github.com/ahaostudy/onlinejudge/app/api/handler/user"
	"github.com/ahaostudy/onlinejudge/app/api/mw/jwt"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRoute(r *server.Hertz) {
	api := r.Group("/api/v1")

	apiUser := api.Group("/user")
	{
		apiUser.POST("/register", user.Register)
		apiUser.POST("/login", jwt.Middleware.LoginHandler)
		apiUser.GET("/refresh", jwt.Middleware.RefreshHandler)
		apiUser.POST("/captcha", user.GetCaptcha)

		apiUser.Use(jwt.Middleware.MiddlewareFunc())

		apiUser.GET("", user.GetUser)
		apiUser.GET("/:id", user.GetUserById)
		apiUser.GET("/un/:username", user.GetUserByUsername)
		apiUser.POST("", user.CreateUser)
		apiUser.PUT("/:id", user.UpdateUser)
	}
}
