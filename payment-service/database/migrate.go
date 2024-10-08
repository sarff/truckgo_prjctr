package database

import (
	"github.com/alexandear/truckgo/payment-service/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Payment{},
	)
}
