package test

import (
	"testing"

	"github.com/shaikhjunaidx/pennywise-backend/internal/budget"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/shaikhjunaidx/pennywise-backend/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupBudgetService() (*budget.BudgetService, *mocks.MockBudgetRepository) {
	mockRepo := new(mocks.MockBudgetRepository)
	service := budget.NewBudgetService(mockRepo)
	return service, mockRepo
}

func TestBudgetService_CreateBudget(t *testing.T) {
	service, mockRepo := setupBudgetService()

	budget := &models.Budget{
		UserID:          1,
		CategoryID:      nil, // Assuming this is an overall budget
		AmountLimit:     1000.0,
		SpentAmount:     0,
		RemainingAmount: 1000.0,
		BudgetMonth:     "09",
		BudgetYear:      2024,
	}

	mockRepo.On("Create", budget).Return(nil)

	result, err := service.CreateBudget(budget.UserID, budget.CategoryID, budget.AmountLimit, budget.BudgetMonth, budget.BudgetYear)

	assert.NoError(t, err)
	assert.Equal(t, budget, result)
	assert.Equal(t, 1000.0, result.RemainingAmount)
	assert.Equal(t, 0.0, result.SpentAmount)

	mockRepo.AssertExpectations(t)
}

func TestBudgetService_UpdateBudget(t *testing.T) {
	service, mockRepo := setupBudgetService()

	budgetID := uint(1)
	newAmountLimit := 2000.0

	existingBudget := &models.Budget{
		ID:              budgetID,
		UserID:          1,
		CategoryID:      nil,
		AmountLimit:     1000.0,
		SpentAmount:     500.0,
		RemainingAmount: 500.0,
		BudgetMonth:     "09",
		BudgetYear:      2024,
	}

	mockRepo.On("FindByID", budgetID).Return(existingBudget, nil)
	mockRepo.On("Update", mock.Anything).Return(nil)

	result, err := service.UpdateBudget(budgetID, newAmountLimit)

	assert.NoError(t, err)
	assert.Equal(t, newAmountLimit, result.AmountLimit)
	assert.Equal(t, newAmountLimit-existingBudget.SpentAmount, result.RemainingAmount)

	mockRepo.AssertExpectations(t)
}

func TestBudgetService_AddTransactionToBudget(t *testing.T) {
	service, mockRepo := setupBudgetService()

	userID := uint(1)
	categoryID := uint(1)
	transactionAmount := 200.0
	month := "09"
	year := 2024

	existingBudget := &models.Budget{
		ID:              1,
		UserID:          userID,
		CategoryID:      &categoryID,
		AmountLimit:     1000.0,
		SpentAmount:     300.0,
		RemainingAmount: 700.0,
		BudgetMonth:     month,
		BudgetYear:      year,
	}

	expectedBudget := &models.Budget{
		SpentAmount:     500.0,
		RemainingAmount: 500.0,
	}

	mockRepo.On("FindByUserIDAndCategoryID", userID, &categoryID, month, year).Return(existingBudget, nil)
	mockRepo.On("Update", mock.Anything).Return(nil)

	result, err := service.AddTransactionToBudget(userID, &categoryID, transactionAmount, month, year)

	assert.NoError(t, err)
	assert.Equal(t, expectedBudget.SpentAmount, result.SpentAmount)
	assert.Equal(t, expectedBudget.RemainingAmount, result.RemainingAmount)

	mockRepo.AssertExpectations(t)
}
