package models

import "github.com/alexandear/truckgo/internal/auth/models"

type Payment struct {
	ID         int         `gorm:"primary_key"`
	NumberCard int         `gorm:"unique;not null"` // Додати індекс
	UserID     uint        `gorm:"not null"`
	User       models.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
