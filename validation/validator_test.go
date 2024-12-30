package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

type UserDto struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name" validate:"required,min=2,max=100,msg=First name is required and must be between 2 and 100 characters"`
	LastName  string `json:"last_name" validate:"required,min=2,max=100,msg=Last name is required and must be between 2 and 100 characters"`
	Email     string `json:"email" validate:"required,email,msg=A valid email is required"`
	Password  string `json:"password" validate:"required,Password,msg=Password must be 8-20 characters long and contain at least one uppercase letter; one lowercase letter; one digit and one special character"`
}

func TestValidate_failed(t *testing.T) {
	user := UserDto{
		Id:        1,
		FirstName: "",
		LastName:  "F",
		Email:     "testemail.com",
		Password:  "password",
	}

	validationRequest := New(user,
		map[string]func(fl validator.FieldLevel) bool{"Password": ValidatePassword()})

	errors := validationRequest.Validate()

	// assert that all the fields failed
	if len(errors) != 4 {
		t.Errorf("Expected 4 errors but got %d", len(errors))
	}

	if errors["FirstName"] != "First name is required and must be between 2 and 100 characters" {
		t.Errorf("Expected 'First name is required and must be between 2 and 100 characters' but got %s", errors["FirstName"])
	}

	if errors["LastName"] != "Last name is required and must be between 2 and 100 characters" {
		t.Errorf("Expected 'Last name is required and must be between 2 and 100 characters' but got %s", errors["LastName"])
		fmt.Println(errors["LastName"])
	}

	if errors["Email"] != "A valid email is required" {
		t.Errorf("A valid email is required' but got %s", errors["Email"])
	}

	if errors["Password"] != "Password must be 8-20 characters long and contain at least one uppercase letter; one lowercase letter; one digit and one special character" {
		t.Errorf("Expected 'Password must be 8-20 characters long and contain at least one uppercase letter; one lowercase letter; one digit and one special character' but got %s", errors["Password"])
	}
}

func TestValidate_successful(t *testing.T) {
	user := UserDto{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "test@email.com",
		Password:  "Password101$",
	}

	validationRequest := New(user,
		map[string]func(fl validator.FieldLevel) bool{"Password": ValidatePassword()})

	errors := validationRequest.Validate()

	// assert that errors is empty
	if len(errors) != 0 {
		t.Errorf("Expected 0 errors but got %d", len(errors))
	}
}
