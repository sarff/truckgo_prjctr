package repository

import (
	"github.com/alexandear/truckgo/order-service/internal/models"
	"gorm.io/gorm"
)

// TODO пошук одного по айді, пошук одного по критерію (чи є таке в го?) як мінімум по потрібному
// TODO пошук багатьох з фільтрацією ?
type Repository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll(page int, limit int) ([]models.Order, int64, error) {
	offset := (page - 1) * limit
	var orders []models.Order
	var total int64

	err := r.db.
		Model(&models.Order{}).
		Count(&total).
		Limit(limit).
		Offset(offset).
		Find(&orders).Error

	return orders, total, err
}

func (r *Repository) FindOneByID(id int) (models.Order, error) {
	var order models.Order

	err := r.db.
		Model(&models.Order{}).
		First(&order, id).Error

	return order, err
}
