package stockentry

import (
	"time"

	"github.com/rdurbina/stocky/internal/product"
	"github.com/rdurbina/stocky/internal/user"
)

type Model struct {
	ID       int
	Date     time.Time
	Product  product.Model
	Employee user.Model
	Amount   int
}
