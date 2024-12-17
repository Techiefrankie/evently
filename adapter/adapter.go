package adapter

import (
	"evently/api"
	"evently/models"
	"evently/security"
	"fmt"
	"time"
)

func EventDtoToModel(dto api.EventDto) *models.Event {
	return &models.Event{
		Name:        dto.Name,
		Description: dto.Description,
		Location:    dto.Location,
		DateTime:    time.Now(),
		UserId:      dto.UserId,
	}
}

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

func UserDtoToModel(dto api.UserDto) models.User {
	encryptedPassword, err := security.GetEncryptedPassword(dto.Password)
	if err != nil {
		fmt.Println("Error encrypting password: ", err)
		return models.User{}
	}

	return models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  encryptedPassword,
	}
}
