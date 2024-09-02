package dtos

type ProductsDTO struct {
	Total    int           `json:"total"`
	Products []*ProductDTO `json:"products"`
}

type ProductDTO struct {
	ID string `json:"id"`
	NewProduct
}

type ProductUpdates struct {
	NewProduct
}

type NewProduct struct {
	Code        string  `json:"code" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description" binding:"required"`
}
