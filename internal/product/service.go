package product

import (
	"github.com/go-playground/validator/v10"
	"github.com/rdurbina/stocky/internal/brand"
	"github.com/rdurbina/stocky/internal/category"
	apperror "github.com/rdurbina/stocky/pkg/app-error"
)

/*
The service for the product model is in charge of validating requests, save data and
return response DTOs.
*/

type Service struct {
	repository         Repository
	brandRepository    brand.Repository
	categoryRepository category.Repository
}

// Validates a product creation request and saves it using the repository
func (s Service) Save(request CreateRequest) (Response, *apperror.ServiceError) {
	var requestErrors []error
	validationErrors := apperror.NewValidationError()
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		requestErrors = TranslateErrors(err.(validator.ValidationErrors))
		validationErrors.AppendErr(requestErrors...)
		return Response{}, validationErrors
	}

	category, categoryErr := s.categoryRepository.FindById(request.CategoryID)
	brand, brandErr := s.brandRepository.FindById(request.BrandID)

	if categoryErr != nil || brandErr != nil {
		validationErrors.AppendErr(categoryErr, brandErr)
		return Response{}, validationErrors
	}

	productModel := ToModel(request, category, brand)
	persistedProduct, err := s.repository.Save(productModel)

	if err != nil {
		return Response{}, apperror.NewInternalError(err)
	}

	return ToResponse(persistedProduct), nil

}

// Creates an instance of the Service struct
func NewService(r Repository, cr category.Repository, br brand.Repository) Service {
	return Service{
		repository:         r,
		brandRepository:    br,
		categoryRepository: cr,
	}
}
