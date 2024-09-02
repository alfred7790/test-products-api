package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"test-products-api/domain/repositories"
	"test-products-api/domain/services"
	"test-products-api/infrastructure/database/sqlite"
	"test-products-api/infrastructure/web/controllers"
)

type Server struct {
	router       *gin.Engine
	db           *gorm.DB
	repositories *Repositories
	services     *Services
	controllers  *Controllers
}

type Repositories struct {
	ProductRepo repositories.Product
}

type Services struct {
	ProductServ *services.ProductService
}

type Controllers struct {
	ProductCtrl *controllers.ProductController
}

func NewServer() *Server {
	s := &Server{}
	s.db = sqlite.NewSQLDB()
	s.repositories = &Repositories{ProductRepo: sqlite.NewProductRepository(s.db)}
	s.services = &Services{ProductServ: services.NewProductService(s.repositories.ProductRepo)}
	s.controllers = &Controllers{ProductCtrl: controllers.NewProductController(*s.services.ProductServ)}
	s.router = NewGinRouter(s.controllers)
	return s
}

func (s *Server) Run() {
	err := s.router.Run(":1010")
	if err != nil {
		panic(err.Error())
	}
}
