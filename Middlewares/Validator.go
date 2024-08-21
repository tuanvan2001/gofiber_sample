package Middlewares

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Param       string
		Value       interface{}
	}

	XValidator struct {
		validator *validator.Validate
	}
)

var validate = validator.New()

func (v XValidator) Validate(data interface{}, messages map[string]string) (map[string]string, bool) {
	validationErrors := make(map[string]string)
	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrors[err.Field()] = getMessage(err, messages)
		}
		return validationErrors, true
	}
	return nil, false
}

func getMessage(fieldError validator.FieldError, messages map[string]string) string {
	key := fmt.Sprintf("%s.%s", fieldError.Field(), fieldError.Tag())
	if msg, exists := messages[key]; exists {
		return msg
	}
	return fmt.Sprintf("Trường %s không hợp lệ.", fieldError.Field())
}

var Validator = &XValidator{validator: validate}
