package brand

import "github.com/rdurbina/stocky/internal/abstractions"

type Repository interface {
	abstractions.BaseRepository[Model]
}
