package mocks

import (
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/stretchr/testify/mock"
)

type MockBudgetRepository struct {
	mock.Mock
}

func (m *MockBudgetRepository) Create(budget *models.Budget) error {
	args := m.Called(budget)
	return args.Error(0)
}

func (m *MockBudgetRepository) Update(budget *models.Budget) error {
	args := m.Called(budget)
	return args.Error(0)
}

func (m *MockBudgetRepository) DeleteByID(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBudgetRepository) FindByID(id uint) (*models.Budget, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Budget), args.Error(1)
}

func (m *MockBudgetRepository) FindAllByUserID(userID uint) ([]*models.Budget, error) {
	args := m.Called(userID)
	return args.Get(0).([]*models.Budget), args.Error(1)
}

func (m *MockBudgetRepository) FindByUserIDAndCategoryID(userID uint, categoryID *uint, month string, year int) (*models.Budget, error) {
	args := m.Called(userID, categoryID, month, year)
	return args.Get(0).(*models.Budget), args.Error(1)
}

func (m *MockBudgetRepository) FindAllByUserIDAndMonthYear(userID uint, month string, year int) ([]*models.Budget, error) {
	args := m.Called(userID, month, year)
	return args.Get(0).([]*models.Budget), args.Error(1)
}
