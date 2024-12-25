package validation

import (
	"errors"
	"github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
	"log"
)

// Request is a struct to hold the request and regex validations
type Request struct {
	Request interface{}
	Regex   map[string]string
}

// New is a function to create a new validation request
func New(request interface{}, regex map[string]string) Request {
	return Request{
		Request: request,
		Regex:   regex,
	}
}

// Validate is a function to validate request
func (req Request) Validate() map[string]string {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validationErrors := make(map[string]string)

	// register regex validations for the respective fields
	var field string
	if len(req.Regex) > 0 {
		for functionName, regex := range req.Regex {
			err := validate.RegisterValidation(functionName, func(fl validator.FieldLevel) bool {

				field = fl.FieldName()
				re := regexp2.MustCompile(regex, regexp2.None)
				match, er := re.MatchString(fl.Field().String())

				if er != nil {
					log.Println(er)
					return false
				}

				return match
			})

			if err != nil {
				// add the field and error to the map
				validationErrors[field] = err.Error()
			}
		}
	}

	err := validate.Struct(req.Request)

	if err != nil {
		// Check if the error is of type validator.ValidationErrors
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			for _, fieldError := range errs {
				// Add the field name and validation error to the map
				validationErrors[fieldError.Field()] = fieldError.Error()
			}
		}
	}

	return validationErrors
}
