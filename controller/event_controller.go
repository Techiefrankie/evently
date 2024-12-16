package controller

import (
	"evently/adapter"
	"evently/api"
	"evently/models"
	"evently/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetEvent(context *gin.Context) {
	id := context.Param("id")
	eventId, err := strconv.Atoi(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, getResponse("Invalid event ID", http.StatusBadRequest))
		return
	}

	fmt.Println("Getting event with ID: ", eventId)

	event := services.GetEvent(uint(eventId))

	if *event != (models.Event{}) {
		context.JSON(http.StatusOK, adapter.EventToDto(*event))
	} else {
		context.JSON(http.StatusNotFound, getResponse("Event not found", http.StatusNotFound))
	}
}

func GetEvents(context *gin.Context) {
	fmt.Println("Getting all events")
	events := services.GetEvents()

	if events != nil {
		context.JSON(http.StatusOK, adapter.EventsToDtos(events))
		return
	} else {
		context.JSON(http.StatusNotFound, getResponse("No events found", http.StatusNotFound))
	}
}

func CreateEvent(context *gin.Context) {
	var event api.EventDto
	err := context.BindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, getResponse(err.Error(), http.StatusBadRequest))
		return
	}

	fmt.Println("Creating event: ", event)

	// save event
	err = services.SaveEvent(adapter.EventDtoToModel(event))

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, getResponse("Failed to create event", http.StatusInternalServerError))
		return
	}

	// return response
	context.JSON(http.StatusCreated, getResponse("Event created", http.StatusCreated))
}

func DeleteEvent(context *gin.Context) {
	id := context.Param("id")
	eventId, err := strconv.Atoi(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, getResponse("Invalid event ID", http.StatusBadRequest))
		return
	}

	fmt.Println("Deleting event with ID: ", eventId)

	err = services.DeleteEvent(uint(eventId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, getResponse("Failed to delete event", http.StatusInternalServerError))
		return
	}

	context.JSON(http.StatusOK, getResponse("Event deleted", http.StatusOK))
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, getResponse("Welcome to Evently!", http.StatusOK))
}

func getResponse(msg string, code int) api.Response {
	return api.Response{
		Message:    msg,
		StatusCode: code,
	}
}
