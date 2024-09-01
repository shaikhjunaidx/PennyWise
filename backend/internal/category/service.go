package category

import (
	"errors"

	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
	"github.com/shaikhjunaidx/pennywise-backend/models"
)

type CategoryService struct {
	Repo        CategoryRepository
	UserService *user.UserService
}

func NewCategoryService(repo CategoryRepository, userService *user.UserService) *CategoryService {
	return &CategoryService{
		Repo:        repo,
		UserService: userService,
	}
}

func (s *CategoryService) AddCategory(username, name, description string) (*models.Category, error) {
	user, err := s.UserService.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	category := &models.Category{
		UserID:      user.ID,
		Name:        name,
		Description: description,
	}

	if err := s.Repo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) GetCategoryByID(username string, id uint) (*models.Category, error) {
	user, err := s.UserService.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	category, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if category.UserID != user.ID {
		return nil, errors.New("access denied: category does not belong to the user")
	}

	return category, nil
}

func (s *CategoryService) GetAllCategories(username string) ([]*models.Category, error) {
	user, err := s.UserService.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	categories, err := s.Repo.FindAllByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryService) UpdateCategory(username string, id uint, name, description string) (*models.Category, error) {
	user, err := s.UserService.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	category, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if category.UserID != user.ID {
		return nil, errors.New("access denied: category does not belong to the user")
	}

	category.Name = name
	category.Description = description

	if err := s.Repo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) DeleteCategory(username string, id uint) error {
	user, err := s.UserService.FindByUsername(username)
	if err != nil {
		return err
	}

	category, err := s.Repo.FindByID(id)
	if err != nil {
		return err
	}

	if category.UserID != user.ID {
		return errors.New("access denied: category does not belong to the user")
	}

	if err := s.Repo.DeleteByID(id); err != nil {
		return err
	}

	return nil
}
