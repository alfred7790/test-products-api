package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-products-api/domain/services"
	"test-products-api/infrastructure/dtos"
	"test-products-api/infrastructure/filters"
	"test-products-api/infrastructure/mappers"
	"test-products-api/infrastructure/utils"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(s services.ProductService) *ProductController {
	return &ProductController{productService: s}
}

func (ctrl *ProductController) NewProduct(c *gin.Context) {
	var newProduct dtos.NewProduct

	err := c.ShouldBindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewAppError(http.StatusBadRequest, "Invalid new product payload", err.Error()))
		return
	}

	product := mappers.ToNewProductModel(&newProduct)
	err = ctrl.productService.CreateProduct(product)
	if err != nil {
		c.JSON(utils.StatusCode(err), err)
		return
	}

	productDTO := mappers.ToProductDTO(product)
	c.JSON(http.StatusCreated, productDTO)
}

func (ctrl *ProductController) ProductByID(c *gin.Context) {
	id := c.Param("id")

	product, err := ctrl.productService.GetProduct(id)
	if err != nil {
		c.JSON(utils.StatusCode(err), err)
		return
	}

	productDTO := mappers.ToProductDTO(product)
	c.JSON(http.StatusOK, productDTO)
}

func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var productUpdates dtos.ProductUpdates
	err := c.ShouldBindJSON(&productUpdates)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewAppError(http.StatusBadRequest, "Invalid product updates payload", err.Error()))
		return
	}

	product := mappers.ToProductUpdates(id, &productUpdates)
	err = ctrl.productService.UpdateProduct(product)
	if err != nil {
		c.JSON(utils.StatusCode(err), err)
		return
	}

	c.JSON(http.StatusOK, "Product updated")
}

func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.productService.DeleteProduct(id)
	if err != nil {
		c.JSON(utils.StatusCode(err), err)
		return
	}

	c.JSON(http.StatusOK, "Product deleted")
}

func (ctrl *ProductController) FindProducts(c *gin.Context) {
	var filtersParams filters.ProductFilter

	err := c.ShouldBindQuery(&filtersParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewAppError(http.StatusBadRequest, "Invalid params for filters", err.Error()))
		return
	}

	total, products, err := ctrl.productService.FilterProducts(&filtersParams)
	if err != nil {
		c.JSON(utils.StatusCode(err), err)
		return
	}

	response := &dtos.ProductsDTO{
		Total:    total,
		Products: mappers.ToProductsModel(products),
	}

	c.JSON(http.StatusOK, response)
}
