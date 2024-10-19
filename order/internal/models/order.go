package models

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Status int

const (
	StatusNew        Status = iota // when order was created
	StatusAccepted                 // when order in status new so it can be accepted
	StatusInProgress               // when driver started order processing and status is accepted
	StatusDone                     // when driver finished it - можна тільки з in progress
	StatusCancelled                // when order was incorrect - можна тільки customer з new
)

// TODO think about string status representation
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

func ValidateStatus(order Order, newStatus Status) error {
	currentStatus := order.Status
	err := status.Errorf(codes.FailedPrecondition, "Cannot change %d order status to %d", newStatus, currentStatus)

	switch newStatus {
	case StatusAccepted:
		if currentStatus != StatusNew {
			return err
		}
	case StatusInProgress:
		if currentStatus != StatusAccepted {
			return err
		}
	case StatusCancelled:
		if currentStatus != StatusAccepted {
			return err
		}
	case StatusDone:
		if currentStatus != StatusInProgress {
			return err
		}
	default:
		return err
	}

	return nil
}
