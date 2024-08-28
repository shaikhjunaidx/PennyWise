package models

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Username     string    `json:"username" gorm:"not null;unique"`
	Email        string    `json:"email" gorm:"not null;unique"`
	PasswordHash string    `json:"-" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
