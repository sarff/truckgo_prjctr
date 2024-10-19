package models

import (
	"github.com/google/uuid"
)

type Status int

const (
	StatusNew        Status = iota // when created
	StatusAccepted                 // when accepted - assign driver id
	StatusInProgress               // when driver started to do it - можна тільки з accepted
	StatusDone                     // when driver finished it - можна тільки з in progress
	StatusCancelled                // when order was incorrect - можна тільки customer з new
)

type Order struct {
	ID         uint    `gorm:"primary_key;auto_increment"`
	Number     string  `gorm:"unique;not null;index"`
	Status     Status  `gorm:"not null;default:0"`
	Price      float64 `gorm:"not null;default:0"`
	UserID     uint    `gorm:"not null"`
	DriverID   uint    `gorm:"null"`
	IsArchived bool    `gorm:"default:false"`
}

func NewOrder(price float64, userID uint) *Order {
	return &Order{
		Number: generateOrderNumber(),
		Status: StatusNew,
		Price:  price,
		UserID: userID,
	}
}

func generateOrderNumber() string {
	return uuid.New().String()
}
