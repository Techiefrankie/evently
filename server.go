package main

import (
	"evently/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() {
	server := gin.Default()

	// server routes
	server.GET("/", controller.Home)
	server.POST("event/create", controller.CreateEvent)
	server.GET("event/", controller.GetEvents)
	server.GET("event/:id", controller.GetEvent)
	server.DELETE("event/delete/:id", controller.DeleteEvent)

	err := server.Run()

	if err != nil {
		fmt.Println(err)
		return
	}
}