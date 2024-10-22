package repository

import (
	"github.com/alexandear/truckgo/payment/internal/models"
	"gorm.io/gorm"
)

type Payment struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *Payment {
	return &Payment{db: db}
}

func (r *Payment) CreateSuccessful(amount uint32, currency string, externalUID uint32, pmID uint32,
	userID uint32,
) (uint32, error) {
	return r.create(amount, currency, externalUID, models.StatusSuccess, pmID, userID)
}

func (r *Payment) CreateFail(amount uint32, currency string, externalUID uint32, pmID uint32,
	userID uint32,
) (uint32, error) {
	return r.create(amount, currency, externalUID, models.StatusFail, pmID, userID)
}

func (r *Payment) create(amount uint32, currency string, externalUID uint32, status models.Status, pmID uint32,
	userID uint32,
) (uint32, error) {
	payment := models.NewPayment(amount, currency, externalUID, status, pmID, userID)
	result := r.db.Create(&payment)

	return payment.ID, result.Error
}

func (r *Payment) Update(payment models.Payment, updates map[string]any) error {
	result := r.db.Model(&payment).Where("id = ?", payment.ID).Updates(updates)
	return result.Error
}

func (r *Payment) FindAllByUser(page int, limit int, filters map[string]any) ([]models.Payment, int64, error) {
	offset := (page - 1) * limit
	var payments []models.Payment
	var total int64

	query := r.db.Model(&models.Payment{})
	for column, value := range filters {
		query = query.Where(column+" = ?", value)
	}

	err := query.Count(&total).
		Limit(limit).
		Offset(offset).
		Find(&payments).Error

	return payments, total, err
}

func (r *Payment) FindOneByID(id uint) (models.Payment, error) {
	var payment models.Payment

	err := r.db.
		Model(&models.Payment{}).
		First(&payment, id).Error

	return payment, err
}
