package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func New() (*validator.Validate, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("json")
	})

	err := validate.RegisterValidation("phoneNumber", phoneNumberValidator)
	if err != nil {
		return nil, err
	}

	return validate, nil
}
