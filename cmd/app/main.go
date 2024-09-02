package main

import (
	"test-products-api/infrastructure/web/router"
)

func main() {
	s := router.NewServer()
	s.Run()
}
