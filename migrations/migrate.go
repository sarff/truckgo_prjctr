package migrations

import (
	auth "github.com/alexandear/truckgo/internal/auth/models"
	orders "github.com/alexandear/truckgo/internal/orders/models"
	payments "github.com/alexandear/truckgo/internal/payments/models"
	shippings "github.com/alexandear/truckgo/internal/shippings/models"
	"github.com/alexandear/truckgo/pkg/logging"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	logging.InitLogger()
	err := db.AutoMigrate(
		&auth.User{}, &auth.TypeUser{}, &auth.Driver{}, &auth.Client{}, // users
		&orders.Order{},
		&payments.Payment{},
		&shippings.Shipping{},
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
