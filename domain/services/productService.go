package services

import (
	"test-products-api/domain/models"
	"test-products-api/domain/repositories"
	"test-products-api/infrastructure/filters"
)

type ProductService struct {
	productRepository repositories.Product
}

func NewProductService(r repositories.Product) *ProductService {
	return &ProductService{productRepository: r}
}

func (s *ProductService) FilterProducts(filters *filters.ProductFilter) (int, []*models.Product, error) {
	if filters.Page == 0 {
		filters.Page = 1
	}

	if filters.Limit == 0 {
		filters.Limit = 30
	}

	return s.productRepository.Filter(filters)
}
func (s *ProductService) CreateProduct(p *models.Product) error {
	return s.productRepository.Create(p)
}
func (s *ProductService) UpdateProduct(p *models.Product) error {
	product, err := s.productRepository.Get(p.ID)
	if err != nil {
		return err
	}

	if p.Name != "" && product.Name != p.Name {
		product.Name = p.Name
	}

	if p.Code != "" && product.Code != p.Code {
		product.Code = p.Code
	}

	if p.Description != "" && product.Description != p.Description {
		product.Description = p.Description
	}

	if p.Price != 0 && product.Price != p.Price {
		product.Price = p.Price
	}

	return s.productRepository.Update(product)
}
func (s *ProductService) GetProduct(id string) (*models.Product, error) {
	return s.productRepository.Get(id)
}
func (s *ProductService) DeleteProduct(id string) error {
	_, err := s.productRepository.Get(id)
	if err != nil {
		return err
	}
	return s.productRepository.Delete(id)
}
