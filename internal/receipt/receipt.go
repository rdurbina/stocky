package receipt

import (
	"time"

	"github.com/rdurbina/stocky/internal/product"
	"github.com/rdurbina/stocky/internal/user"
)

type Model struct {
	ID       int
	Product  product.Model
	Amount   int
	Date     time.Time
	Total    float64
	Employee user.Model
}
