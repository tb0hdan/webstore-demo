package memory

import (
	"sync"

	"webstore-demo/internal/server/api"
	"webstore-demo/pkg/utils"
	"webstore-demo/pkg/xerrors"
)

type Store struct {
	lock     *sync.Mutex
	products []api.Product
	sales    []api.Sale
}

func (s *Store) GetProducts() ([]api.Product, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.products, nil
}

func (s *Store) AddProduct(product api.Product) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if utils.Index(s.products, func(p api.Product) bool {
		return p.Id == product.Id
	}) != -1 {
		return xerrors.ErrProductAlreadyExists
	}
	s.products = append(s.products, product)
	return nil
}

func (s *Store) AddSale(sale api.Sale) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.sales = append(s.sales, sale)
	return nil
}

func NewStore() *Store {
	return &Store{
		products: []api.Product{},
		sales:    []api.Sale{},
		lock:     &sync.Mutex{},
	}
}
