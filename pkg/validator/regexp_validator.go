package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func phoneNumberValidator(fl validator.FieldLevel) bool {
	pattern := regexp.MustCompile(`\+380\d{9}`)

	return pattern.MatchString(fl.Field().String())
}
