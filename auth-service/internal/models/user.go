package models

import (
	"time"
)

// Чисто приклад - треба змінити під потреби сервісу
type User struct {
	ID         uint      `gorm:"primary_key"`
	Login      string    `gorm:"unique;not null"` // email
	Password   string    `gorm:"not null"`
	FullName   string    `gorm:"not null"`
	TypeUserID uint      `gorm:"not null"`
	Status     bool      `gorm:"not null;default:true"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoCreateTime"`
}

// client or driver
type TypeUser struct {
	ID   uint   `gorm:"primary_key"`
	Type string `gorm:"not null"`
}

type Driver struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	License   string  `gorm:"not null"`
	CarModel  string  `gorm:"not null"`
	CarNumber string  `gorm:"not null"`
	Rating    float64 `gorm:"default:0"`
}

type Client struct {
	ID      uint   `gorm:"primaryKey"`
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Phone   string `gorm:"not null"`
	Address string `gorm:"not null"`
}
