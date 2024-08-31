package category

import (
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{DB: db}
}

func (r *CategoryRepositoryImpl) Create(category *models.Category) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepositoryImpl) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := r.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) FindAll() ([]*models.Category, error) {
	var categories []*models.Category
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepositoryImpl) Update(category *models.Category) error {
	return r.DB.Save(category).Error
}

func (r *CategoryRepositoryImpl) DeleteByID(id uint) error {
	return r.DB.Delete(&models.Category{}, id).Error
}

func (r *CategoryRepositoryImpl) FindByName(name string) (*models.Category, error) {
	var category models.Category
	if err := r.DB.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
