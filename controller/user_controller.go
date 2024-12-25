package controller

import (
	"evently/adapter"
	"evently/api"
	"evently/services"
	"evently/utils"
	"evently/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(context *gin.Context) {
	var userRequest api.UserDto

	err := context.BindJSON(&userRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse("Invalid request"+err.Error(), http.StatusBadRequest))
		return
	}

	// validate request
	validationRequest := validation.New(userRequest,
		map[string]string{"password": `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,20}$`}, nil)

	validationErrors := validationRequest.Validate()
	if len(validationErrors) > 0 {
		context.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	err = services.SaveUser(adapter.UserDtoToModel(userRequest))
	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.GetResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	context.JSON(http.StatusOK, utils.GetResponse("User created successfully", http.StatusOK))
}

func Login(context *gin.Context) {
	var loginRequest api.LoginDto

	err := context.BindJSON(&loginRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse("Invalid request"+err.Error(), http.StatusBadRequest))
		return
	}

	// validate request
	validationRequest := validation.Request{Request: loginRequest}
	validationErrors := validationRequest.Validate()

	if len(validationErrors) > 0 {
		context.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	response, err := services.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, utils.GetResponse("Invalid email or password", http.StatusUnauthorized))
		return
	}

	context.JSON(http.StatusOK, response)
}
