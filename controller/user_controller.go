package controller

import (
	"evently/adapter"
	"evently/api"
	"evently/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(context *gin.Context) {
	var userRequest api.UserDto

	err := context.BindJSON(&userRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, getResponse("Invalid request"+err.Error(), http.StatusBadRequest))
		return
	}

	err = services.SaveUser(adapter.UserDtoToModel(userRequest))
	if err != nil {
		context.JSON(http.StatusInternalServerError, getResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	context.JSON(http.StatusOK, getResponse("User created successfully", http.StatusOK))
}
