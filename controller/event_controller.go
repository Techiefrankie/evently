package controller

import (
	"evently/adapter"
	"evently/api"
	"evently/models"
	"evently/services"
	"evently/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetEvent(context *gin.Context) {
	id := context.Param("id")
	eventId, err := strconv.Atoi(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse("Invalid event ID", http.StatusBadRequest))
		return
	}

	fmt.Println("Getting event with ID: ", eventId)

	event := services.GetEvent(uint(eventId))

	if *event != (models.Event{}) {
		context.JSON(http.StatusOK, adapter.EventToDto(*event))
	} else {
		context.JSON(http.StatusNotFound, utils.GetResponse("Event not found", http.StatusNotFound))
	}
}

func GetEvents(context *gin.Context) {
	fmt.Println("Getting all events")
	events := services.GetEvents()

	if events != nil {
		context.JSON(http.StatusOK, adapter.EventsToDtos(events))
		return
	} else {
		context.JSON(http.StatusNotFound, utils.GetResponse("No events found", http.StatusNotFound))
	}
}

func CreateEvent(context *gin.Context) {
	authValue, _ := context.Get("auth")
	authResponse, _ := authValue.(api.AuthResponse)

	var event api.EventDto
	err := context.BindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse(err.Error(), http.StatusBadRequest))
		return
	}

	// bind the id of the authenticated user
	event.UserId = authResponse.UserId

	fmt.Println("Creating event: ", event)

	// save event
	err = services.SaveEvent(adapter.EventDtoToModel(event))

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, utils.GetResponse("Failed to create event", http.StatusInternalServerError))
		return
	}

	// return response
	context.JSON(http.StatusCreated, utils.GetResponse("Event created", http.StatusCreated))
}

func DeleteEvent(context *gin.Context) {
	id := context.Param("id")
	eventId, err := strconv.Atoi(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse("Invalid event ID", http.StatusBadRequest))
		return
	}

	fmt.Println("Deleting event with ID: ", eventId)

	userId := context.MustGet("auth").(api.AuthResponse).UserId
	err = services.DeleteEvent(uint(eventId), userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.GetResponse("Failed to delete event", http.StatusInternalServerError))
		return
	}

	context.JSON(http.StatusOK, utils.GetResponse("Event deleted", http.StatusOK))
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, utils.GetResponse("Welcome to Evently!", http.StatusOK))
}
