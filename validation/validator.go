package validation

import (
	"errors"
	"github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"strings"
)

type Request struct {
	Body interface{}
	Func map[string]func(fl validator.FieldLevel) bool
}

func New(body interface{}, Func map[string]func(fl validator.FieldLevel) bool) Request {
	return Request{Body: body, Func: Func}
}

func (req Request) Validate() map[string]string {
	validate := validator.New()
	validationErrors := make(map[string]string)

	//  Register custom validation for the 'msg' tag for providing custom error messages
	validate.RegisterValidation("msg", func(fl validator.FieldLevel) bool { return true })

	for tag, fn := range req.Func {
		validate.RegisterValidation(tag, fn)
	}

	if err := validate.Struct(req.Body); err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			for _, fieldError := range errs {
				fieldType := reflect.TypeOf(req.Body)
				if fieldType.Kind() == reflect.Ptr {
					fieldType = fieldType.Elem() // Dereference the pointer to get the actual type
				}

				field, _ := fieldType.FieldByName(fieldError.StructField())
				if msg := getErrorMessage(field); msg != "" {
					validationErrors[fieldError.Field()] = msg
				} else {
					panic("Error message not set for field: " + fieldError.Field())
				}
			}
		}
	}

	return validationErrors
}

func getErrorMessage(field reflect.StructField) string {
	for _, part := range strings.Split(field.Tag.Get("validate"), ",") {
		if kv := strings.Split(part, "="); kv[0] == "msg" {
			return kv[1]
		}
	}
	return ""
}

func ValidatePassword() func(fl validator.FieldLevel) bool {
	regex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,20}$`

	return func(fl validator.FieldLevel) bool {
		re := regexp2.MustCompile(regex, regexp2.None)
		match, er := re.MatchString(fl.Field().String())

		if er != nil {
			log.Println(er)
			return false
		}

		return match
	}
}
