package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func NewValidationError(err validator.ValidationErrors) []string {
	fieldErrors := make([]string, 0, len(err))

	for _, e := range err {
		var msg string

		switch e.Tag() {
			case "required":
				msg = "%s is required"
			case "email":
				msg = "%s must be a valid email"
			case "min":
				msg = "%s must have at least %s characters"
			case "max":
				msg = "%s must have at most %s characters"
			default:
				msg = "failed to validate %s: " + e.Tag()
		}

		if e.Param() != "" {
			fieldErrors = append(fieldErrors, fmt.Sprintf(msg, e.Field(), e.Param()))
		} else {
			fieldErrors = append(fieldErrors, fmt.Sprintf(msg, e.Field()))
		}
	}

	return fieldErrors
}

func ValidateBody[T any](body *T) ([]string, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(body)

	if err != nil {
		var invalidValidationErr *validator.InvalidValidationError
		if errors.As(err, &invalidValidationErr) {
			return nil, err
		}

		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			return NewValidationError(validationErrs), nil
		}
	}

	return nil, nil
}
