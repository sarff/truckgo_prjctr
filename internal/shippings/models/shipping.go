package models

import (
	"github.com/alexandear/truckgo/internal/orders/models"
	"time"
)

type Shipping struct {
	ID             uint         `gorm:"primary_key"`
	AddressStart   string       `gorm:"not null"`
	AddressFinish  string       `gorm:"not null"`
	ShipPriceCents int64        `gorm:"not null;default:0"`
	OrderID        uint         `gorm:"not null"`
	Order          models.Order `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CreatedAt      time.Time    `gorm:"autoCreateTime"`
	UpdatedAt      time.Time    `gorm:"autoCreateTime"`
}
