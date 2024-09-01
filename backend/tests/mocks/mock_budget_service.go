package mocks

import (
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/stretchr/testify/mock"
)

type MockBudgetService struct {
	mock.Mock
}

func (m *MockBudgetService) CreateBudget(username string, categoryID *uint, amountLimit float64, month string, year int) (*models.Budget, error) {
	args := m.Called(username, categoryID, amountLimit, month, year)
	return args.Get(0).(*models.Budget), args.Error(1)
}

func (m *MockBudgetService) UpdateBudget(budgetID uint, amountLimit float64) (*models.Budget, error) {
	args := m.Called(budgetID, amountLimit)
	return args.Get(0).(*models.Budget), args.Error(1)
}

func (m *MockBudgetService) DeleteBudget(budgetID uint) error {
	args := m.Called(budgetID)
	return args.Error(0)
}

func (m *MockBudgetService) GetBudgetByID(budgetID uint) (*models.Budget, error) {
	args := m.Called(budgetID)
	return args.Get(0).(*models.Budget), args.Error(1)
}

func (m *MockBudgetService) GetBudgetsForUser(username string) ([]*models.Budget, error) {
	args := m.Called(username)
	return args.Get(0).([]*models.Budget), args.Error(1)
}

func (m *MockBudgetService) GetBudgetForUserAndCategory(username string, categoryID *uint, month string, year int) (*models.Budget, error) {
	args := m.Called(username, categoryID, month, year)
	return args.Get(0).(*models.Budget), args.Error(1)
}

func (m *MockBudgetService) AddTransactionToBudget(userID uint, categoryID *uint, transactionAmount float64, month string, year int) (*models.Budget, error) {
	args := m.Called(userID, categoryID, transactionAmount, month, year)
	return args.Get(0).(*models.Budget), args.Error(1)
}

func (m *MockBudgetService) CalculateOverallBudget(username string) (*models.Budget, error) {
	args := m.Called(username)
	return args.Get(0).(*models.Budget), args.Error(1)
}
