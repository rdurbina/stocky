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
