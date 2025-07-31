package product

import (
	"errors"

	"github.com/rdurbina/stocky/internal/abstractions"
)

type Repository interface {
	abstractions.BaseRepository[Model]
}

/*
In memory version of the product repository, it uses a map to keep track of product records,
use only for development purposes
*/
type InMemoryRepository struct {
	productCounter int
	products       map[int]Model
}

/*InMemoryRepository methods*/

// Retrives a product record from memory
func (in InMemoryRepository) FindById(id int) (Model, error) {
	value, ok := in.products[id]
	if !ok {
		return Model{}, errors.New("product not found")
	}
	return value, nil
}

// Saves a product record into volatile memory
func (in *InMemoryRepository) Save(product Model) (Model, error) {
	product.ID = in.productCounter
	in.products[product.ID] = product
	in.productCounter++
	return product, nil
}

// Updates a product record saved in volatile memory
func (in InMemoryRepository) Update(id int, product Model) (Model, error) {
	_, ok := in.products[id]
	if !ok {
		return Model{}, errors.New("product not found")
	}
	in.products[id] = product
	return product, nil

}

// Deletes a product from volatile memory
func (in InMemoryRepository) Delete(id int) error {
	_, ok := in.products[id]
	if !ok {
		return errors.New("product not found")
	}
	return nil
}

// Creates an instace of an in memory product repository, use only for development purposes
func NewInMemoryRepository() Repository {
	return &InMemoryRepository{
		productCounter: 0,
		products:       make(map[int]Model),
	}
}
