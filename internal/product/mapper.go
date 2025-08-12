package product

import (
	"github.com/rdurbina/stocky/internal/brand"
	"github.com/rdurbina/stocky/internal/category"
)

// Maps a Product Model into a Response DTO, it'll include all the fields
func ToResponse(product Model) Response {
	return Response{
		ID:          product.ID,
		Name:        product.Name,
		Brand:       product.Brand.Name,
		Category:    product.Category.Name,
		Sku:         product.Sku,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Available:   product.Available,
		Hidden:      product.Hidden,
	}

}

/*
Maps a creation request into a model, it will set the following fields with default values:

- Stock: Set to 0
- Available: Set to false
- Hidden: Set to false.
*/
func ToModel(request CreateRequest, category category.Model, brand brand.Model) Model {
	return Model{
		Name:        request.Name,
		Brand:       brand,
		Category:    category,
		Description: request.Description,
		Price:       request.Price,
		Stock:       0,
		Available:   false,
		Hidden:      false,
	}
}
