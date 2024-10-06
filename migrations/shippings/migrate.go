package shippings

import (
	shippings "github.com/alexandear/truckgo/internal/shippings/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&shippings.Shipping{},
	)
}
