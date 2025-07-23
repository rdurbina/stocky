package product

import (
	"github.com/go-playground/validator/v10"
)

// Translates struct level validation errors
// into a user friendly version of them.
func TranslateErrors(fieldErrors []validator.FieldError) map[string]string {
	requestErrors := make(map[string]string)
	for _, fieldErr := range fieldErrors {
		e := fieldErr.Field()
		switch e {
		case "Name":
			requestErrors[e] = "The name should be between 2 and 50 characters long."
		case "BrandID":
			requestErrors[e] = "The brand ID is required."
		case "CategoryID":
			requestErrors[e] = "The category ID is required."
		case "Description":
			requestErrors[e] = "The description should be between 2 and 100 characters long"
		case "Price":
			requestErrors[e] = "The price cannot be negative"
		}
	}
	return requestErrors
}
