package services

import (
	"errors"
	"evently/api"
	"evently/config"
	"evently/models"
	"evently/security"
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

func DeleteEvent(id uint, userId int) error {
	var db = *config.GetDbInstance()
	var event = GetEvent(id)

	if event != nil && event.UserId == userId {
		err := db.Delete(&event)

		if err != nil {
			return err.Error
		}
	} else {
		return errors.New("event not found")
	}

	return nil
}

func Login(email string, password string) (api.AuthResponse, error) {
	var db = *config.GetDbInstance()
	var user models.User
	db.Where("email = ?", email).First(&user)

	if user != (models.User{}) {

		if security.PasswordMatches(password, user.Password) {
			// generate token
			token, err := security.GenerateToken(user)

			if err != nil {
				return api.AuthResponse{}, err
			}

			return api.AuthResponse{
				Email:       user.Email,
				UserId:      user.Id,
				AccessToken: token,
				ExpiresIn:   30,
				Roles:       []string{"USER"},
			}, nil
		} else {
			return api.AuthResponse{}, errors.New("password is incorrect")
		}
	}

	return api.AuthResponse{}, errors.New(fmt.Sprintf("email %v is incorrect", email))
}
