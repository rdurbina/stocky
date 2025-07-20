package abstractions

type BaseRepository[T any] interface {
	FindById(id int) (T, error)
	Save(entity T) (T, error)
	Update(id int, entity T) (T, error)
	Delete(id int) error
}
