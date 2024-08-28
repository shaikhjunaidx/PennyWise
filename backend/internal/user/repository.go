package user

import (
	"github.com/shaikhjunaidx/pennywise-backend/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Update(user *models.User) error
	Delete(user *models.User) error
}
