package models

import (
	"time"
)

type Auth struct {
	ID        uint32    `gorm:"primary_key"`
	Login     string    `gorm:"unique;not null"` // email
	Password  string    `gorm:"not null"`
	TypeUser  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}
