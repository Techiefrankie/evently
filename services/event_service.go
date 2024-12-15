package services

import (
	"errors"
	"evently/config"
	"evently/models"
	"time"
)

func SaveEvent(event *models.Event) error {
	var db = *config.GetDbInstance()
	event.DateTime = time.Now()
	event.UserId = 1

	err := db.Create(&event)

	if err != nil {
		return err.Error
	}

	return nil
}

func GetEvents() []models.Event {
	var db = *config.GetDbInstance()
	var events []models.Event
	db.Find(&events)

	return events
}

func GetEvent(id uint) *models.Event {
	var db = *config.GetDbInstance()
	var event models.Event
	db.First(&event, id)

	return &event
}

func DeleteEvent(id uint) error {
	var db = *config.GetDbInstance()
	var event = GetEvent(id)

	if event != nil {
		err := db.Delete(&event)

		if err != nil {
			return err.Error
		}
	} else {
		return errors.New("event not found")
	}

	return nil
}
