package models

import "time"

type Transaction struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id" gorm:"not null"`
	User            User      `json:"-" gorm:"foreignKey:UserID"`
	CategoryID      uint      `json:"category_id"`
	Category        Category  `json:"-" gorm:"foreignKey:CategoryID"`
	Amount          float64   `json:"amount" gorm:"not null"`
	Description     string    `json:"description,omitempty"`
	TransactionDate time.Time `json:"transaction_date" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
