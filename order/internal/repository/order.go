package repository

import (
	"github.com/alexandear/truckgo/order/internal/models"
	"gorm.io/gorm"
)

// TODO пошук одного по айді, пошук одного по критерію (чи є таке в го?) як мінімум по потрібному
// TODO пошук багатьох з фільтрацією ?
type Order struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *Order {
	return &Order{db: db}
}

func (r *Order) Create(price float64, userID uint) (uint, error) {
	order := models.NewOrder(price, userID)
	result := r.db.Create(&order)

	return order.ID, result.Error
}

func (r *Order) Update(order models.Order) {
	/// Select with Map
	//// User's ID is `111`:
	//db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	//// UPDATE users SET name='hello' WHERE id=111;
	//
	//db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	//// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
	//
	//// Select with Struct (select zero value fields)
	//db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
	//// UPDATE users SET name='new_name', age=0 WHERE id=111;
	//
	//// Select all fields (select all fields include zero value fields)
	//db.Model(&user).Select("*").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})
	//
	//// Select all fields but omit Role (select all fields include zero value fields)
	//db.Model(&user).Select("*").Omit("Role").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})
}

func (r *Order) UpdateStatus(order models.Order, status models.Status) error {
	err := r.db.Model(&order).Select("status").Updates(map[string]interface{}{"status": status}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Order) FindAll(page int, limit int) ([]models.Order, int64, error) {
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

func (r *Order) FindOneByID(id uint) (models.Order, error) {
	var order models.Order

	err := r.db.
		Model(&models.Order{}).
		First(&order, id).Error

	return order, err
}
