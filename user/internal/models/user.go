package models

import (
	"time"
)

type User struct {
	ID         uint32    `gorm:"primary_key;autoIncrement"`
	Login      string    `gorm:"unique;not null"` // email
	FullName   string    `gorm:"not null"`
	TypeUserID uint32    `gorm:"not null"`
	TypeUser   TypeUser  `gorm:"foreignKey:TypeUserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Status     bool      `gorm:"not null;default:true"`
	Phone      string    `gorm:"not null"`
	Rating     float32   `gorm:"default:0"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoCreateTime"`
	Latitude   float64   `gorm:"default:0"`
	Longitude  float64   `gorm:"default:0"`
}

// TypeUser customer or driver.
type TypeUser struct {
	ID   uint32 `gorm:"primary_key;autoIncrement"`
	Type string `gorm:"not null"`
}

type Driver struct {
	ID        uint32 `gorm:"primaryKey;autoIncrement"`
	UserID    uint32 `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	License   string `gorm:"not null"`
	CarModel  string `gorm:"not null"`
	CarNumber string `gorm:"not null"`
}
