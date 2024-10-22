package database

import (
	"gorm.io/gorm"

	"github.com/alexandear/truckgo/shipping/internal/models"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Shipping{},
	)
}
