package database

import (
	"fmt"
	"github.com/alexandear/truckgo/user/internal/models"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.User{}, &models.TypeUser{}, &models.Driver{}, // users
	); err != nil {
		return err
	}

	var count int64

	db.Model(&models.TypeUser{}).Count(&count)

	// If there are no records, add 'customer' and 'driver'
	if count == 0 {
		users := []models.TypeUser{
			{Type: "customer", ID: 1},
			{Type: "driver", ID: 2},
		}
		if err := db.Create(&users).Error; err != nil {
			return err
		}
	}

	if viper.GetBool("is_demo") {
		var countUser int64
		db.Model(&models.User{}).Count(&countUser)

		if countUser == 0 {
			err := SeedDemoData(db)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func SeedDemoData(db *gorm.DB) error {
	var users []models.User
	var drivers []models.Driver

	// drivers
	for i := 1; i <= 5; i++ {
		drivers = append(drivers, models.Driver{
			ID:        uint32(i),
			UserID:    uint32(i),
			License:   gofakeit.Numerify("########"),
			CarModel:  gofakeit.RandomString([]string{"VW", "Audi", "BMW", "Volvo"}),
			CarNumber: gofakeit.Numerify("АА###МЕ"),
		})
		users = append(users, models.User{
			ID:         uint32(i),
			Login:      fmt.Sprintf("driver%d@example.com", i),
			FullName:   gofakeit.LastName(),
			TypeUserID: 2,
			Status:     true,
			Phone:      gofakeit.Phone(),
			Rating:     gofakeit.Float32Range(0, 100),
		})
	}

	// customers
	for i := 6; i <= 10; i++ {
		users = append(users, models.User{
			ID:         uint32(i),
			Login:      fmt.Sprintf("customer%d@example.com", i),
			FullName:   gofakeit.LastName(),
			TypeUserID: 1,
			Status:     true,
			Phone:      gofakeit.Phone(),
			Rating:     gofakeit.Float32Range(0, 100),
		})
	}

	for _, user := range users {
		db.Create(&user)
	}

	for _, driver := range drivers {
		db.Create(&driver)
	}

	return nil
}
