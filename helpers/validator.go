package helper

import (
	"errors"

	"github.com/agussuartawan/golang-pos/data/response"
	"github.com/go-playground/validator/v10"
)

// Validator validates the given struct and returns an array of
// response.ValidationFailsResponse containing all the validation errors.
//
// This function uses the github.com/go-playground/validator/v10 package
// to validate the struct. The validation errors are then converted to
// response.ValidationFailsResponse and returned as an array.
//
// If there are no validation errors, this function returns nil.
func Validator(t interface{}) []response.ValidationFailsResponse {
	validate := validator.New()
	err := validate.Struct(t)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		out := make([]response.ValidationFailsResponse, len(validationErrors))
		for i, validationError := range validationErrors {
			out[i] = response.ValidationFailsResponse{
				Field:   validationError.Field(),
				Message: tagToMessage(validationError.Tag(), validationError.Param()),
			}
		}
		return out
	}

	return nil
}

// tagToMessage returns a human-readable message for the given
// validation tag. The message is based on the tag and the param.
// The returned message is a string.
//
// If the tag is not recognized, this function returns an empty string.
func tagToMessage(tag string, param string) string {
	switch tag {
	case "required":
		return "Field is required"
	case "max":
		return "Field is too long, max " + param
	case "min":
		return "Field is too short, min " + param
	case "email":
		return "Field is not a valid email"
	case "gte":
		return "Field is too low, min " + param
	case "lte":
		return "Field is too high, max " + param
	case "oneof":
		return "Field must be one of " + param
	case "gt":
		return "Field must be greater than " + param
	case "lt":
		return "Field must be less than " + param
	default:
		return ""
	}
}