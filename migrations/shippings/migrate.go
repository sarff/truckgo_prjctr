package shippings

import (
	shippings "github.com/alexandear/truckgo/internal/shippings/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&shippings.Shipping{},
	)

	if err != nil {
		return err
	}
	return nil
}
