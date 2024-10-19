package database

import (
	"github.com/alexandear/truckgo/auth/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.Auth{}, // auth
	); err != nil {
		return err
	}
	return nil
}
