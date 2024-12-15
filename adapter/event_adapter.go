package adapter

import (
	"evently/api"
	"evently/models"
)

func EventToDto(event models.Event) api.EventDto {
	return api.EventDto{
		ID:          event.Id,
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		DateTime:    event.DateTime.Format("Jan 02, 2006 at 3:04:05 PM"),
		UserId:      event.UserId,
	}
}

func EventsToDtos(events []models.Event) []api.EventDto {
	var eventDtos []api.EventDto

	for _, event := range events {
		eventDtos = append(eventDtos, EventToDto(event))
	}

	return eventDtos
}
