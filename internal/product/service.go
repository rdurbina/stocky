package product

import (
	"github.com/go-playground/validator/v10"
)

type Service struct {
	Repository Repository
}

func (s Service) Save(request CreateRequest) (Response, map[string]string) {
	var requestErrors map[string]string
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		requestErrors = TranslateErrors(err.(validator.ValidationErrors))
		return Response{}, requestErrors
	}
	//TODO
	return Response{}, requestErrors
}

func New() Service {
	return Service{}
}
