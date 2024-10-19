package models

type Status int

const (
	StatusNew        Status = iota // when created
	StatusAccepted                 // when accepted - assign driver id
	StatusInProgress               // when driver started to do it - можна тільки з accepted
	StatusDone                     // when driver finished it - можна тільки з in progress
	StatusCancelled                // when order was incorrect - можна тільки customer з new
)

type Order struct {
	ID         int    `gorm:"primary_key"`
	Number     int    `gorm:"unique;not null"` // Додати індекс
	Status     Status `gorm:"not null;default:0"`
	Price      int64  `gorm:"not null;default:0"` // cents
	UserID     uint   `gorm:"not null"`
	DriverId   uint   `gorm:"null"`
	IsArchived bool   `gorm:"default:false"`
}

func NewOrder(price int64, userID uint) *Order {
	return &Order{Status: StatusNew, Price: price, UserID: userID}
}

//TODO create getters and setters
//TODO
