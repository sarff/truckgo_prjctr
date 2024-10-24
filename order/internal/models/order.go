package models

import (
	"github.com/google/uuid"
)

type Status int32

const (
	StatusNew        Status = iota // when order was created
	StatusAccepted                 // when order in status new so it can be accepted
	StatusInProgress               // when driver started order processing and status is accepted
	StatusDone                     // when driver finished it - можна тільки з in progress
	StatusCancelled                // when order was incorrect - можна тільки customer з new
)

type Order struct {
	ID         uint32  `gorm:"primary_key;auto_increment"`
	Number     string  `gorm:"unique;not null;index"`
	Status     Status  `gorm:"not null;default:0"`
	Price      float64 `gorm:"not null;default:0"`
	UserID     uint32  `gorm:"not null"`
	DriverID   uint32  `gorm:"null"`
	IsArchived bool    `gorm:"default:false"`
}

func NewOrder(price float64, userID uint32) *Order {
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

func ValidateStatus(order Order, newStatus Status) bool {
	currentStatus := order.Status

	switch newStatus {
	case StatusAccepted:
		if currentStatus != StatusNew {
			return false
		}
	case StatusInProgress:
		if currentStatus != StatusAccepted {
			return false
		}
	case StatusCancelled:
		if currentStatus != StatusAccepted {
			return false
		}
	case StatusDone:
		if currentStatus != StatusInProgress {
			return false
		}
	case StatusNew:
		if currentStatus != StatusAccepted {
			return false
		}
	default:
		return false
	}

	return true
}
