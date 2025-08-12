package apperror

// Error types
const (
	INTERNAL_ERROR = iota
	INVALID_REQUEST
)

/*
ServiceError is an error wrapper, it collects all the validation errors from a service
or it can hold a short-circuiting error on it's internalError field, it also contains a Code variable
*/
type ServiceError struct {
	validationErrors []error
	internalError    error
	Code             int
}

// Appends an error to the error slice, it will skip nil values
func (se *ServiceError) AppendErr(errs ...error) *ServiceError {
	for _, err := range errs {
		if err == nil {
			continue
		}
		se.validationErrors = append(se.validationErrors, err)
	}
	return se
}

/*
Depending on the error code, it calls the Error() method for each error from the validationErrors slice
or calls Error() from the internalError field.
*/
func (se ServiceError) Error() []string {
	if se.Code == INVALID_REQUEST {
		var errMessages []string
		for _, err := range se.validationErrors {
			errMessages = append(errMessages, err.Error())
		}
		return errMessages
	} else {
		return []string{se.internalError.Error()}
	}

}

// Returns and instance of ServiceError prepared to hold validation errors
func NewValidationError() *ServiceError {
	return &ServiceError{
		Code:          INVALID_REQUEST,
		internalError: nil,
	}
}

// Wraps an error instance into an instance of server error (typically a short-circuiting error)
func NewInternalError(err error) *ServiceError {
	return &ServiceError{
		Code:             INTERNAL_ERROR,
		validationErrors: nil,
	}
}
