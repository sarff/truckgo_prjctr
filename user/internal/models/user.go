package models

import (
	"time"
)

type User struct {
	ID         uint      `gorm:"primary_key"`
	Login      string    `gorm:"unique;not null"` // email
	FullName   string    `gorm:"not null"`
	TypeUserID uint      `gorm:"not null"`
	TypeUser   TypeUser  `gorm:"foreignKey:TypeUserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Status     bool      `gorm:"not null;default:true"`
	Phone      string    `gorm:"not null"`
	Rating     float64   `gorm:"default:0"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoCreateTime"`
}

// TypeUser customer or driver.
type TypeUser struct {
	ID   uint   `gorm:"primary_key"`
	Type string `gorm:"not null"`
}

type Driver struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	License   string `gorm:"not null"`
	CarModel  string `gorm:"not null"`
	CarNumber string `gorm:"not null"`
	//	Position  string  `gorm:"not null"`
}
