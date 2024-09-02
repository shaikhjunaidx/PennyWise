package mocks

import (
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/stretchr/testify/mock"
)

type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) AddCategory(username, name, description string) (*models.Category, error) {
	args := m.Called(username, name, description)
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockCategoryService) DeleteCategory(username string, id uint) error {
	args := m.Called(username, id)
	return args.Error(0)
}
