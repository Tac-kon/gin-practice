// products_service.go

package services

import (
	"github.com/gouser/api/domain/products"
	"github.com/gouser/api/utils/errors"
)

// CreateProduct - Service
func CreateProduct(product products.Product) (*products.Product, *errors.ApiErr) {
	if err := product.Validate(); err != nil {
		return nil, err
	}

	if err := product.Save(); err != nil {
		return nil, err
	}

	return &product, nil
}

// GetProduct - Service
func GetProduct(productID uint64) (*products.Product, *errors.ApiErr) {
	p := &products.Product{ID: productID}
	if err := p.Get(); err != nil {
		return nil, err
	}
	return p, nil
}