package database

import (
	"github.com/alexandear/truckgo/user-service/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.User{}, &models.TypeUser{}, &models.Driver{}, &models.Client{}, // users
	); err != nil {
		return err
	}
	var count int64

	db.Model(&models.TypeUser{}).Count(&count)

	// If there are no records, add 'customer' and 'driver'
	if count == 0 {
		users := []models.TypeUser{
			{Type: "customer"},
			{Type: "driver"},
		}

		if err := db.Create(&users).Error; err != nil {
			return err
		}
	}
	return nil
}
