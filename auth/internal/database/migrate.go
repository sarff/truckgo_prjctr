package database

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/alexandear/truckgo/auth/internal/models"
	"github.com/alexandear/truckgo/auth/internal/services"
)

func Migrate(ctx context.Context, db *gorm.DB) error {
	if err := db.WithContext(ctx).AutoMigrate(
		&models.Auth{}, // auth
	); err != nil {
		return err
	}

	if viper.GetBool("is_demo") {
		var count int64
		db.WithContext(ctx).Model(&models.Auth{}).Count(&count)

		if count == 0 {
			err := SeedDemoData(ctx, db)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SeedDemoData(ctx context.Context, db *gorm.DB) error {
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
		db.WithContext(ctx).Create(&user)
	}

	return nil
}
