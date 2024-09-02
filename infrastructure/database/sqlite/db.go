package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"test-products-api/domain/models"
)

func NewSQLDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
