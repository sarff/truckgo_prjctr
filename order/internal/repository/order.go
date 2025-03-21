package repository

import (
	"gorm.io/gorm"

	"github.com/alexandear/truckgo/order/internal/models"
)

type Order struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *Order {
	return &Order{db: db}
}

func (r *Order) Create(price float64, userID uint32) (uint32, error) {
	order := models.NewOrder(price, userID)
	result := r.db.Create(&order)

	return order.ID, result.Error
}

func (r *Order) Update(order models.Order, updates map[string]any) error {
	result := r.db.Model(&order).Where("id = ?", order.ID).Updates(updates)
	return result.Error
}

func (r *Order) FindAll(page int, limit int, filters map[string]any) ([]models.Order, int64, error) {
	offset := (page - 1) * limit
	var orders []models.Order
	var total int64

	query := r.db.Model(&models.Order{})
	for column, value := range filters {
		query = query.Where(column+" = ?", value)
	}

	err := query.Count(&total).
		Limit(limit).
		Offset(offset).
		Find(&orders).Error

	return orders, total, err
}

func (r *Order) FindOneByID(id uint32) (models.Order, error) {
	var order models.Order

	err := r.db.
		Model(&models.Order{}).
		First(&order, id).Error

	return order, err
}
