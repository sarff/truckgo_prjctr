package models

type Payment struct {
	ID         int  `gorm:"primary_key"`
	NumberCard int  `gorm:"unique;not null"` // Додати індекс
	UserID     uint `gorm:"not null"`
}
