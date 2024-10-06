package payments

import (
	payments "github.com/alexandear/truckgo/internal/payments/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&payments.Payment{},
	)
}
