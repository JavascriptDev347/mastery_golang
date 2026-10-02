package task

import (
	"errors"
	"sync"
)

var ErrProductNotFound = errors.New("product not found")
var ErrInsufficientStock = errors.New("insufficient stock")

type InventoryManager interface {
	AddProduct(id string, name string, stock int, price float64) error
	GetProduct(id string) (*Product, error)
	UpdateStock(id string, quantity int) error
	DeleteProduct(id string) error
}

type Product struct {
	ID    string
	Name  string
	Stock int
	Price float64
}

type storeInventory struct {
	sync.RWMutex
	products map[string]*Product
}

func NewStoreInventory() *storeInventory {
	return &storeInventory{
		products: make(map[string]*Product),
	}
}

func (s *storeInventory) AddProduct(id string, name string, stock int, price float64) error {
	//mutex
	s.Lock()
	defer s.Unlock()

	if _, exists := s.products[id]; exists {
		return errors.New("product already exists")
	}

	s.products[id] = &Product{
		ID:    id,
		Name:  name,
		Stock: stock,
		Price: price,
	}
	return nil
}

func (s *storeInventory) GetProduct(id string) (Product, error) {
	//mutex for read. the main reason for using RLock is to allow concurrent reads without blocking
	s.RLock()
	defer s.RUnlock()

	product, ok := s.products[id]
	if !ok {
		return Product{}, ErrProductNotFound
	}
	return *product, nil
}

func (s *storeInventory) UpdateStock(id string, quantity int) error {
	//mutex
	s.Lock()
	defer s.Unlock()

	if product, exists := s.products[id]; exists {
		product.Stock += quantity
		return nil
	}
	return ErrProductNotFound
}

func (s *storeInventory) DeleteProduct(id string) error {
	//mutex
	s.Lock()
	defer s.Unlock()

	if _, exists := s.products[id]; exists {
		delete(s.products, id)
		return nil
	}
	return ErrProductNotFound
}

func main() {}
