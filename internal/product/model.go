package product

import (
	brand "github.com/rdurbina/stocky/internal/Brand"
	"github.com/rdurbina/stocky/internal/category"
)

type Model struct {
	ID          int
	Name        string
	Brand       brand.Model
	Category    category.Model
	Sku         string
	Description string
	Price       float64
	Stock       int
}

type CreateRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=50"`
	BrandID     int     `json:"brandId" validate:"required"`
	CategoryID  int     `json:"categoryId" validate:"required"`
	Description string  `json:"description" validate:"required,min=2,max=100"`
	Price       float64 `json:"price" validate:"required,min=0"`
}

type Response struct {
	Name        string
	ID          int
	Brand       string
	Category    string
	Sku         string
	Description string
	Price       float64
	Stock       int
}
