package models

// Чисто приклад - треба змінити під потреби сервісу
type Order struct {
	ID     int    `gorm:"primary_key"`
	Number int    `gorm:"unique;not null"` // Додати індекс
	Status string `gorm:"not null"`
	Price  int64  `gorm:"not null;default:0"` // cents
	UserID uint   `gorm:"not null"`
}
