package customvalidator

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func CustomValidaton(validation map[string]func(fl validator.FieldLevel) bool) echo.Validator {
	v := validator.New()

	for name, function := range validation {
		v.RegisterValidation(name, function)
	}

	return &CustomValidator{validator: v}
}

func ValidateDateFormat(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	if dateStr == "" || fl.Field().IsZero() {
		return true
	}

	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
