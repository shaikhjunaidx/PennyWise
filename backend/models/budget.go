package models

import "time"

type Budget struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id" gorm:"not null"`
	User            User      `json:"-" gorm:"foreignKey:UserID"`
	CategoryID      *uint     `json:"category_id,omitempty"`
	Category        Category  `json:"-" gorm:"foreignKey:CategoryID"`
	AmountLimit     float64   `json:"amount_limit" gorm:"not null"`
	SpentAmount     float64   `json:"spent_amount" gorm:"not null"`
	RemainingAmount float64   `json:"remaining_amount" gorm:"not null"`
	BudgetMonth     string    `json:"budget_month" gorm:"size:2;not null"`
	BudgetYear      int       `json:"budget_year" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
