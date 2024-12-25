package main

import (
	"evently/controller"
	"evently/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func RegisterRoutes() {
	server := gin.Default()
	err := server.SetTrustedProxies([]string{"127.0.0.1"})

	if err != nil {
		log.Println(err)
		return
	}

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

	err = server.Run(":8080")

	if err != nil {
		log.Println(err)
		return
	}
}
