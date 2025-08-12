package product

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

/*
Translates struct level validation errors
into a user friendly version of them.
*/

func TranslateErrors(fieldErrors []validator.FieldError) []error {
	var requestErrors []error
	var err error
	for _, fieldErr := range fieldErrors {
		e := fieldErr.Field()
		switch e {
		case "Name":
			err = errors.New("the name should be between 2 and 50 characters long")
			requestErrors = append(requestErrors, err)
		case "BrandID":
			err = errors.New("the brand ID is required")
			requestErrors = append(requestErrors, err)
		case "CategoryID":
			err = errors.New("the category ID is required")
			requestErrors = append(requestErrors, err)
		case "Description":
			err = errors.New("the description should be between 2 and 100 characters long")
			requestErrors = append(requestErrors, err)
		case "Price":
			err = errors.New("the price cannot be negative")
			requestErrors = append(requestErrors, err)
		}
	}
	return requestErrors
}
