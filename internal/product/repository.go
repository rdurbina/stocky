package product

import (
	"github.com/rdurbina/stocky/internal/abstractions"
	"github.com/rdurbina/stocky/internal/user"
)

type Repository interface {
	abstractions.BaseRepository[user.Model]
}
