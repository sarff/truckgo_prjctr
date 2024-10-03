package payments

import (
	payments "github.com/alexandear/truckgo/internal/payments/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&payments.Payment{},
	)

	if err != nil {
		return err
	}
	return nil
}
