package models

type PaymentMethod struct {
	ID     uint32 `gorm:"primary_key;auto_increment"`
	UserID uint32 `gorm:"not null"`
	Number string `gorm:"not null"`
	Expiry string `gorm:"not null"`
	Cvv    string `gorm:"not null"`
	Name   string `gorm:"not null"`
	Active bool   `gorm:"default:true"`
}

func NewPaymentMethod(userID uint32, number string, expiry string, cvv string, name string) *PaymentMethod {
	return &PaymentMethod{
		UserID: userID,
		Number: number,
		Expiry: expiry,
		Cvv:    cvv,
		Name:   name,
		Active: true,
	}
}
