package types

import "webstore-demo/internal/server/api"

type Store interface {
	// This is a placeholder for the store
	GetProducts() ([]api.Product, error)
	AddProduct(product api.Product) error
	AddSale(sale api.Sale) error
}
