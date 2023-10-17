package config

import (
	"github.com/go-playground/validator/v10"

	validatorHelper "github.com/aldisaputra17/go-micro/validator"
)

// NewValidator.
func NewValidator() *validatorHelper.Validator {
	return &validatorHelper.Validator{Validator: validator.New()}
}
