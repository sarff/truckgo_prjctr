package models

import (
	"time"
)

type Auth struct {
	ID        uint      `gorm:"primary_key"`
	Login     string    `gorm:"unique;not null"` // email
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}
