package repositories

import (
	"test-products-api/domain/models"
	"test-products-api/infrastructure/filters"
)

type Product interface {
	Filter(filters *filters.ProductFilter) (int, []*models.Product, error)
	Create(p *models.Product) error
	Update(p *models.Product) error
	Get(id string) (*models.Product, error)
	Delete(id string) error
}
