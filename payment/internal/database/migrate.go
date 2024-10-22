package database

import (
	"github.com/alexandear/truckgo/payment/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Payment{},
		&models.PaymentMethod{},
	)
}
