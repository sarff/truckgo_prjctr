package database

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/alexandear/truckgo/user/internal/models"
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
	var i uint32
	for i = 1; i <= 5; i++ {
		lat, _ := gofakeit.LatitudeInRange(50.150001, 50.700001)
		long, _ := gofakeit.LongitudeInRange(30.100001, 30.990001)
		drivers = append(drivers, models.Driver{
			ID:        i,
			UserID:    i,
			License:   gofakeit.Numerify("########"),
			CarModel:  gofakeit.RandomString([]string{"VW", "Audi", "BMW", "Volvo"}),
			CarNumber: gofakeit.Numerify("АА###МЕ"),
		})
		users = append(users, models.User{
			ID:         i,
			Login:      fmt.Sprintf("driver%d@example.com", i),
			FullName:   gofakeit.LastName(),
			TypeUserID: 2,
			Status:     true,
			Phone:      gofakeit.Phone(),
			Rating:     gofakeit.Float32Range(0, 100),
			Latitude:   lat,
			Longitude:  long,
		})
	}

	// customers
	for i = 6; i <= 10; i++ {
		lat, _ := gofakeit.LatitudeInRange(50.150001, 50.700001)
		long, _ := gofakeit.LongitudeInRange(30.100001, 30.990001)
		users = append(users, models.User{
			ID:         i,
			Login:      fmt.Sprintf("customer%d@example.com", i),
			FullName:   gofakeit.LastName(),
			TypeUserID: 1,
			Status:     true,
			Phone:      gofakeit.Phone(),
			Rating:     gofakeit.Float32Range(0, 100),
			Latitude:   lat,
			Longitude:  long,
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
