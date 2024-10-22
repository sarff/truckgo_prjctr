package models

type Status int

const (
	StatusNew Status = iota
	StatusSuccess
	StatusFail
	StatusRefunded
)

type Payment struct {
	ID              uint32 `gorm:"primary_key;auto_increment"`
	ExternalID      uint32 `gorm:"not null"`
	PaymentMethodID uint32 `gorm:"not null"`
	UserID          uint32 `gorm:"not null"`
	Amount          uint32 `gorm:"not null"`
	Currency        string `gorm:"not null"`
	Status          Status `gorm:"not null"`
}

func NewPayment(amount uint32, currency string, externalID uint32, status Status, pmID uint32, userID uint32) *Payment {
	return &Payment{
		ExternalID:      externalID,
		PaymentMethodID: pmID,
		Amount:          amount,
		Currency:        currency,
		UserID:          userID,
		Status:          status,
	}
}
