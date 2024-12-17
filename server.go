package main

import (
	"evently/controller"
	"evently/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() {
	server := gin.Default()

	// configure protected routes
	authenticated := server.Group("/", middleware.Authenticate)
	authenticated.POST("event/create", middleware.Authenticate, controller.CreateEvent)
	authenticated.GET("event/", controller.GetEvents)
	authenticated.GET("event/:id", controller.GetEvent)
	authenticated.DELETE("event/delete/:id", controller.DeleteEvent)

	// server routes
	server.GET("/", controller.Home)
	server.POST("user/create", controller.CreateUser)
	server.POST("user/login", controller.Login)

	err := server.Run()

	if err != nil {
		fmt.Println(err)
		return
	}
}
