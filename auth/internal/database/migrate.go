package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/alexandear/truckgo/auth/internal/models"
	"github.com/alexandear/truckgo/auth/internal/services"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.Auth{}, // auth
	); err != nil {
		return err
	}

	if viper.GetBool("is_demo") {
		var count int64
		db.Model(&models.Auth{}).Count(&count)

		if count == 0 {
			err := SeedDemoData(db)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SeedDemoData(db *gorm.DB) error {
	pass, err := services.HashPassword("password123")
	if err != nil {
		return err
	}

	var users []models.Auth

	// drivers
	var i uint32
	for i = 1; i <= 5; i++ {
		users = append(users, models.Auth{
			Login:    fmt.Sprintf("driver%d@example.com", i),
			Password: pass,
			TypeUser: "driver",
		})
	}

	// customers
	for i = 6; i <= 10; i++ {
		users = append(users, models.Auth{
			Login:    fmt.Sprintf("customer%d@example.com", i),
			Password: pass,
			TypeUser: "customer",
		})
	}

	for _, user := range users {
		db.Create(&user)
	}

	return nil
}
