package mappers

import (
	"strings"
	"test-products-api/domain/models"
	"test-products-api/infrastructure/dtos"
)

func ToProductUpdates(id string, updates *dtos.ProductUpdates) *models.Product {
	return &models.Product{
		ID:          id,
		Code:        strings.TrimSpace(updates.Code),
		Name:        strings.TrimSpace(strings.ToUpper(updates.Name)),
		Price:       updates.Price,
		Description: strings.TrimSpace(updates.Description),
	}
}

func ToNewProductModel(dto *dtos.NewProduct) *models.Product {
	return &models.Product{
		Code:        strings.TrimSpace(dto.Code),
		Name:        strings.TrimSpace(strings.ToUpper(dto.Name)),
		Price:       dto.Price,
		Description: strings.TrimSpace(dto.Description),
	}
}

func ToProductDTO(productModel *models.Product) *dtos.ProductDTO {
	return &dtos.ProductDTO{
		ID: productModel.ID,
		NewProduct: dtos.NewProduct{
			Code:        productModel.Code,
			Name:        productModel.Name,
			Price:       productModel.Price,
			Description: productModel.Description,
		},
	}
}

func ToProductsModel(productsModel []*models.Product) []*dtos.ProductDTO {
	productsDTO := make([]*dtos.ProductDTO, 0)

	for _, productModel := range productsModel {
		productsDTO = append(productsDTO, ToProductDTO(productModel))
	}

	return productsDTO
}
