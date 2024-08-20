package Middlewares

import (
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

func (v XValidator) Validate(data interface{}) (map[string]string, bool) {
	validationErrors := make(map[string]string)
	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrors[err.Field()] = err.Tag() + err.Param()
		}
		return validationErrors, true
	}
	return nil, false
}

var Validator = &XValidator{validator: validate}
