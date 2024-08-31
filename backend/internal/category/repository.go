package category

import "github.com/shaikhjunaidx/pennywise-backend/models"

type CategoryRepository interface {
	Create(category *models.Category) error
	FindByID(id uint) (*models.Category, error)
	FindByName(name string) (*models.Category, error)
	FindAll() ([]*models.Category, error)
	Update(category *models.Category) error
	DeleteByID(id uint) error
}
