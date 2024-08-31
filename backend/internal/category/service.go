package category

import (
	"github.com/shaikhjunaidx/pennywise-backend/models"
)

type CategoryService struct {
	Repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{
		Repo: repo,
	}
}

func (s *CategoryService) AddCategory(name, description string) (*models.Category, error) {
	category := &models.Category{
		Name:        name,
		Description: description,
	}

	if err := s.Repo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	category, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) GetAllCategories() ([]*models.Category, error) {
	categories, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryService) UpdateCategory(id uint, name, description string) (*models.Category, error) {
	category, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	category.Name = name
	category.Description = description

	if err := s.Repo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) DeleteCategory(id uint) error {
	if err := s.Repo.DeleteByID(id); err != nil {
		return err
	}

	return nil
}
