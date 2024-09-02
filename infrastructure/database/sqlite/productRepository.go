package sqlite

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"sync/atomic"
	"test-products-api/domain/models"
	"test-products-api/infrastructure/filters"
	"test-products-api/infrastructure/utils"
	"time"
)

type ProductRepository struct {
	idCounter uint64
	DB        *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Filter(filters *filters.ProductFilter) (int, []*models.Product, error) {
	var (
		total    int64
		products []*models.Product
	)

	query := r.DB.Model(&models.Product{})

	if filters != nil {
		if filters.ID != "" {
			query = query.Where("id = ?", filters.ID)
		}

		if filters.Code != "" {
			query = query.Where("code = ?", filters.Code)
		}

		if filters.Name != "" {
			query = query.Where("name = ?", filters.Name)
		}

		if filters.MinPrice != 0 {
			query = query.Where("price >= ?", filters.MinPrice)
		}

		if filters.MaxPrice != 0 {
			query = query.Where("price <= ?", filters.MaxPrice)
		}
	}

	err := query.Count(&total).Error
	if err != nil {
		return 0, nil, utils.NewAppError(http.StatusInternalServerError, "internal server error trying to query total of products", err.Error())
	}

	err = query.Debug().
		Offset((filters.Page - 1) * filters.Limit).
		Limit(filters.Limit).
		Find(&products).Error
	if err != nil {
		return 0, nil, utils.NewAppError(http.StatusInternalServerError, "internal server error trying to query products", err.Error())
	}

	return int(total), products, nil
}
func (r *ProductRepository) Create(p *models.Product) error {
	p.ID = r.newID()
	p.CreatedAt = time.Now().UTC()
	p.UpdatedAt = p.CreatedAt

	err := r.DB.Create(p).Error
	if err != nil {
		return utils.NewAppError(http.StatusInternalServerError, "Error trying to create a new product", err.Error())
	}
	return nil
}
func (r *ProductRepository) Update(p *models.Product) error {
	p.UpdatedAt = time.Now().UTC()
	err := r.DB.Model(&models.Product{}).
		Where("id = ?", p.ID).
		Updates(p).Error
	if err != nil {
		return utils.NewAppError(http.StatusInternalServerError, "Internal server error trying to update product", err.Error())
	}
	return nil
}
func (r *ProductRepository) Get(id string) (*models.Product, error) {
	var product models.Product
	err := r.DB.Model(&models.Product{}).Where("id = ?", id).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewAppError(http.StatusNotFound, "product doesn't exist", err.Error())
		}
		return nil, utils.NewAppError(http.StatusInternalServerError, "internal server error trying to get product by ID", err.Error())
	}
	return &product, nil
}
func (r *ProductRepository) Delete(id string) error {
	err := r.DB.Unscoped().Delete(&models.Product{ID: id}).Error
	if err != nil {
		return utils.NewAppError(http.StatusInternalServerError, "internal server error trying to delete product", err.Error())
	}
	return nil
}

func (r *ProductRepository) newID() string {
	return fmt.Sprintf("P%d", atomic.AddUint64(&r.idCounter, 1))
}
