package orders

import (
	orders "github.com/alexandear/truckgo/internal/orders/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&orders.Order{},
	)

	if err != nil {
		return err
	}
	return nil
}
