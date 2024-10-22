package models

import (
	"time"
)

type Shipping struct {
	ID             uint      `gorm:"primary_key"`
	AddressStart   string    `gorm:"not null"`
	AddressFinish  string    `gorm:"not null"`
	ShipPriceCents int64     `gorm:"not null;default:0"`
	OrderID        uint      `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoCreateTime"`
}
