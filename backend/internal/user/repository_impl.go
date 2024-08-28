package user

import (
	"errors"

	"github.com/shaikhjunaidx/pennywise-backend/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Create(user *models.User) error {

	if user.Username == "" {
		return errors.New("username cannot be empty")
	}
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	return r.DB.Create(user).Error
}

func (r *UserRepositoryImpl) Update(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*models.User, error) {
	var user models.User

	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) Delete(user *models.User) error {
	return r.DB.Delete(user).Error
}
