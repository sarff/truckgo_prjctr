package repository

import (
	"github.com/alexandear/truckgo/payment/internal/models"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	db *gorm.DB
}

func NewPaymentMethodRepository(db *gorm.DB) *PaymentMethod {
	return &PaymentMethod{db: db}
}

func (r *PaymentMethod) Create(userID uint32, number string, expiry string, cvv string, name string) (uint32, error) {
	paymentMethod := models.NewPaymentMethod(userID, number, expiry, cvv, name)
	result := r.db.Create(&paymentMethod)

	return paymentMethod.ID, result.Error
}

func (r *PaymentMethod) Deactivate(id uint32) error {
	query := r.db.Model(&models.PaymentMethod{})
	result := query.Where("id = ?", id).Update("active", false)

	return result.Error
}

func (r *PaymentMethod) FindOneByUser(userID uint32) (models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod

	query := r.db.Model(&models.PaymentMethod{})
	query = query.Where("user_id = ? AND active = ?", userID, true)

	err := query.Find(&paymentMethod).First(&paymentMethod).Error
	if err != nil {
		return paymentMethod, err
	}

	return paymentMethod, nil
}

func (r *PaymentMethod) FindByID(id uint32) (models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod

	err := r.db.
		Model(&models.PaymentMethod{}).
		First(&paymentMethod, id).Error

	return paymentMethod, err
}
