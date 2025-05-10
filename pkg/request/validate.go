package request

import (
	"github.com/go-playground/validator/v10"
)

func IsValid[T any](loginRequest T) error {
	validate := validator.New()
	err := validate.Struct(loginRequest)
	return err
}
