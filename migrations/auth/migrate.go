package auth

import (
	auth "github.com/alexandear/truckgo/internal/auth/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&auth.User{}, &auth.TypeUser{}, &auth.Driver{}, &auth.Client{}, // users
	)

	if err != nil {
		return err
	}
	var count int64

	db.Model(&auth.TypeUser{}).Count(&count)

	// If there are no records, add 'customer' and 'driver'
	if count == 0 {
		users := []auth.TypeUser{
			{Type: "customer"},
			{Type: "driver"},
		}
		if err = db.Create(&users).Error; err != nil {
			return err
		}
	}
	return nil
}
