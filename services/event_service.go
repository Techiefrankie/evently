package services

import (
	"errors"
	"evently/config"
	"evently/models"
	"fmt"
)

func saveEntity[T any](entity *T) error {
	var db = *config.GetDbInstance()
	err := db.Create(&entity)

	if err != nil {
		return err.Error
	}

	return nil

}

func SaveUser(user models.User) error {
	err := saveEntity(&user)

	if err != nil {
		return err
	}

	return nil
}

func SaveEvent(event *models.Event) error {
	user := GetUser(event.UserId)
	if user != nil {
		event.UserId = user.Id
	} else {
		return errors.New(fmt.Sprintf("user not found with Id: %v", event.UserId))
	}

	err := saveEntity(&event)

	if err != nil {
		return err
	}

	return nil
}

func GetUser(id int) *models.User {
	var db = *config.GetDbInstance()
	var user models.User
	db.First(&user, id) // search for user with id primary key
	// db.First(&user, "first_name = ?", "techie") // search for user with first_name = techie

	return &user
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
