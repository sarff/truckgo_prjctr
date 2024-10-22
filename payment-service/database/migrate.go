package database

import (
	"gorm.io/gorm"

	"github.com/alexandear/truckgo/payment-service/internal/models"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Payment{},
	)
}
