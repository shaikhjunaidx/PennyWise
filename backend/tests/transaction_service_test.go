package test

import (
	"testing"
	"time"

	"github.com/shaikhjunaidx/pennywise-backend/internal/transaction"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/shaikhjunaidx/pennywise-backend/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func setUpTransactionService() (*transaction.TransactionService, *mocks.MockTransactionRepository) {
	mockRepo := new(mocks.MockTransactionRepository)
	service := transaction.NewTransactionService(mockRepo)
	return service, mockRepo
}

func TestTransactionService_AddTransaction(t *testing.T) {
	service, mockRepo := setUpTransactionService()

	transaction := &models.Transaction{
		UserID:          1,
		CategoryID:      1,
		Amount:          100.0,
		Description:     "Groceries",
		TransactionDate: time.Now(),
	}

	mockRepo.On("Create", transaction).Return(nil)

	result, err := service.AddTransaction(transaction.UserID, transaction.CategoryID,
		transaction.Amount, transaction.Description, transaction.TransactionDate)

	assert.NoError(t, err)
	assert.Equal(t, transaction, result)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_UpdateTransaction(t *testing.T) {
	service, mockRepo := setUpTransactionService()

	transaction := &models.Transaction{
		ID:              1,
		UserID:          1,
		CategoryID:      1,
		Amount:          100.0,
		Description:     "Groceries",
		TransactionDate: time.Now(),
	}

	updatedAmount := 200.0
	updatedCategoryID := 2
	updatedDescription := "Updated Groceries"
	updatedTransactionDate := time.Now().AddDate(0, 0, 1)

	mockRepo.On("FindByID", transaction.ID).Return(transaction, nil)
	mockRepo.On("Update", transaction).Return(nil)

	result, err := service.UpdateTransaction(transaction.ID, updatedAmount, uint(updatedCategoryID), updatedDescription, updatedTransactionDate)

	assert.NoError(t, err)
	assert.Equal(t, updatedAmount, result.Amount)
	assert.Equal(t, updatedDescription, result.Description)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_DeleteTransaction(t *testing.T) {
	service, mockRepo := setUpTransactionService()

	transactionID := uint(1)

	mockRepo.On("DeleteByID", transactionID).Return(nil)

	err := service.DeleteTransaction(transactionID)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactionsForUser(t *testing.T) {
	service, mockRepo := setUpTransactionService()

	userID := uint(1)
	transactions := []*models.Transaction{
		{
			UserID:          userID,
			CategoryID:      1,
			Amount:          50.0,
			Description:     "Dinner",
			TransactionDate: time.Now(),
		},
		{
			UserID:          userID,
			CategoryID:      2,
			Amount:          150.0,
			Description:     "Utilities",
			TransactionDate: time.Now(),
		},
	}

	mockRepo.On("FindAllByUserID", userID).Return(transactions, nil)

	result, err := service.GetTransactionsForUser(userID)

	assert.NoError(t, err)
	assert.Equal(t, transactions, result)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactionByID_Success(t *testing.T) {
	service, mockRepo := setUpTransactionService()

	transactionID := uint(1)
	expectedTransaction := &models.Transaction{
		ID:              transactionID,
		UserID:          1,
		CategoryID:      1,
		Amount:          100.0,
		Description:     "Groceries",
		TransactionDate: time.Now(),
	}

	mockRepo.On("FindByID", transactionID).Return(expectedTransaction, nil)

	result, err := service.GetTransactionByID(transactionID)

	assert.NoError(t, err)
	assert.Equal(t, expectedTransaction, result)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactionByID_NotFound(t *testing.T) {
	service, mockRepo := setUpTransactionService()

	transactionID := uint(1)

	mockRepo.On("FindByID", transactionID).Return((*models.Transaction)(nil), assert.AnError)

	result, err := service.GetTransactionByID(transactionID)

	assert.Error(t, err)
	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)
}
