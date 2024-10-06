package orders

import (
	orders "github.com/alexandear/truckgo/internal/orders/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&orders.Order{},
	)
}
