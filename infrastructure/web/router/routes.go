package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewGinRouter(ctrls *Controllers) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running!"})
	})

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Pong!"})
	})

	// V1 Routes
	v1Routes(r, ctrls)
	return r
}

func v1Routes(r *gin.Engine, ctrls *Controllers) {
	v1 := r.Group("v1")
	products := v1.Group("products")
	products.POST("", ctrls.ProductCtrl.NewProduct)
	products.GET("", ctrls.ProductCtrl.FindProducts)
	products.GET(":id", ctrls.ProductCtrl.ProductByID)
	products.PATCH(":id", ctrls.ProductCtrl.UpdateProduct)
	products.DELETE(":id", ctrls.ProductCtrl.DeleteProduct)
}
