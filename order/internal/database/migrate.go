package database

import (
	"strconv"

	"gorm.io/gorm"

	"github.com/alexandear/truckgo/order/internal/models"
	"github.com/spf13/viper"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.Order{},
	); err != nil {
		return err
	}

	if viper.GetBool("is_demo") {
		var countOrder int64
		db.Model(&models.Order{}).Count(&countOrder)

		if countOrder == 0 {
			err := SeedDemoData(db)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func SeedDemoData(db *gorm.DB) error {
	var orders []models.Order

	var i uint32
	for i = 1; i <= 2; i++ {
		orders = append(orders, models.Order{
			ID:         i,
			Number:     strconv.Itoa(int(i)),
			Status:     models.StatusDone,
			Price:      322.22,
			UserID:     6,
			DriverID:   1,
			IsArchived: false,
		})
	}

	for _, order := range orders {
		err := db.Create(&order).Error
		if err != nil {
			return err
		}
	}

	return nil
}
