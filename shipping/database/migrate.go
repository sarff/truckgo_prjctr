package database

import (
	"github.com/alexandear/truckgo/shipping/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Shipping{},
	)
}
