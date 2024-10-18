package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

func ValidateStruct(v interface{}) []string {
	var errorMessages []string
	if err := Validate.Struct(v); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				fieldName := strings.ToLower(strings.TrimPrefix(fieldError.Field(), "Request"))
				switch fieldError.Tag() {
				case "required":
					errorMessages = append(errorMessages, fieldName+" is required")
				case "min":
					minLength := fieldError.Param()
					errorMessages = append(errorMessages, fieldName+" must be at least "+minLength+" characters")
				case "max":
					maxLength := fieldError.Param()
					errorMessages = append(errorMessages, fieldName+" must not exceed "+maxLength+" characters")
				}
			}
		}
	}
	return errorMessages
}
