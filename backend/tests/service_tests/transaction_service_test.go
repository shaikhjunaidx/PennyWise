package test

import (
	"testing"
	"time"

	"github.com/shaikhjunaidx/pennywise-backend/internal/budget"
	"github.com/shaikhjunaidx/pennywise-backend/internal/transaction"
	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/shaikhjunaidx/pennywise-backend/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setUpTransactionService() (*transaction.TransactionService, *mocks.MockTransactionRepository, *mocks.MockUserRepository, *mocks.MockCategoryRepository, *mocks.MockBudgetRepository) {
	mockRepo := new(mocks.MockTransactionRepository)
	mockCategoryRepo := new(mocks.MockCategoryRepository)
	mockUserRepo := &mocks.MockUserRepository{
		Users: make(map[string]*models.User),
	}
	mockBudgetRepo := new(mocks.MockBudgetRepository)

	userService := &user.UserService{Repo: mockUserRepo}
	budgetService := budget.NewBudgetService(mockBudgetRepo, userService)

	service := transaction.NewTransactionService(mockRepo, mockUserRepo, mockCategoryRepo, budgetService)
	return service, mockRepo, mockUserRepo, mockCategoryRepo, mockBudgetRepo
}

func createTestUser(mockUserRepo *mocks.MockUserRepository, username string, id uint) *models.User {
	user := &models.User{
		ID:       id,
		Username: username,
		Email:    username + "@example.com",
	}
	mockUserRepo.Users[username] = user
	return user
}

func createTestTransaction(userID, categoryID uint, amount float64, description string) *models.Transaction {
	return &models.Transaction{
		UserID:          userID,
		CategoryID:      categoryID,
		Amount:          amount,
		Description:     description,
		TransactionDate: time.Now(),
	}
}

func TestTransactionService_AddTransaction(t *testing.T) {
	service, mockRepo, mockUserRepo, _, mockBudgetRepo := setUpTransactionService()

	username := "john_doe"
	user := createTestUser(mockUserRepo, username, 1)
	transaction := createTestTransaction(user.ID, 1, 100.0, "Groceries")

	mockRepo.On("Create", mock.Anything).Return(nil)
	mockBudgetRepo.On("FindByUserIDAndCategoryID", user.ID, &transaction.CategoryID, transaction.TransactionDate.Month().String(), transaction.TransactionDate.Year()).Return(&models.Budget{}, nil)
	mockBudgetRepo.On("Update", mock.AnythingOfType("*models.Budget")).Return(nil)

	result, err := service.AddTransaction(username, transaction.CategoryID, transaction.Amount, transaction.Description, transaction.TransactionDate)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, transaction.Amount, result.Amount)

	mockRepo.AssertExpectations(t)
	mockBudgetRepo.AssertExpectations(t)
}

func TestTransactionService_UpdateTransaction(t *testing.T) {
	service, mockRepo, mockUserRepo, _, mockBudgetRepo := setUpTransactionService()

	username := "john_doe"
	user := createTestUser(mockUserRepo, username, 1)

	transaction := createTestTransaction(user.ID, 2, 100.0, "Groceries")
	transaction.ID = 1

	updatedAmount := 200.0
	updatedCategoryID := uint(2)
	updatedDescription := "Updated Groceries"
	updatedTransactionDate := time.Now().AddDate(0, 0, 1)

	mockRepo.On("FindByID", transaction.ID).Return(transaction, nil)
	mockRepo.On("Update", transaction).Return(nil)
	mockBudgetRepo.On("FindByUserIDAndCategoryID", user.ID, &updatedCategoryID, updatedTransactionDate.Month().String(), updatedTransactionDate.Year()).Return(&models.Budget{}, nil)
	mockBudgetRepo.On("Update", mock.AnythingOfType("*models.Budget")).Return(nil)

	result, err := service.UpdateTransaction(transaction.ID, updatedAmount, updatedCategoryID, updatedDescription, updatedTransactionDate)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updatedAmount, result.Amount)
	assert.Equal(t, updatedDescription, result.Description)

	mockRepo.AssertExpectations(t)
	mockBudgetRepo.AssertExpectations(t)
}

func TestTransactionService_DeleteTransaction(t *testing.T) {
	service, mockRepo, mockUserRepo, _, mockBudgetRepo := setUpTransactionService()

	username := "john_doe"
	user := createTestUser(mockUserRepo, username, 1)
	transactionID := uint(1)
	transaction := createTestTransaction(user.ID, 2, 100.0, "Groceries")
	transaction.ID = transactionID

	mockRepo.On("FindByID", transactionID).Return(transaction, nil)
	mockRepo.On("DeleteByID", transactionID).Return(nil)
	mockBudgetRepo.On("FindByUserIDAndCategoryID", user.ID, &transaction.CategoryID, transaction.TransactionDate.Month().String(), transaction.TransactionDate.Year()).Return(&models.Budget{}, nil)
	mockBudgetRepo.On("Update", mock.AnythingOfType("*models.Budget")).Return(nil)

	err := service.DeleteTransaction(transactionID)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
	mockBudgetRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactionsForUser(t *testing.T) {
	service, mockRepo, mockUserRepo, _, _ := setUpTransactionService()

	username := "john_doe"
	user := createTestUser(mockUserRepo, username, 1)

	transactions := []*models.Transaction{
		createTestTransaction(user.ID, 2, 50.0, "Dinner"),
		createTestTransaction(user.ID, 3, 150.0, "Utilities"),
	}

	mockRepo.On("FindAllByUsername", username).Return(transactions, nil)

	result, err := service.GetTransactionsForUser(username)

	assert.NoError(t, err)
	assert.Equal(t, transactions, result)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactionByID_Success(t *testing.T) {
	service, mockRepo, _, _, _ := setUpTransactionService()

	transactionID := uint(1)
	expectedTransaction := createTestTransaction(1, 2, 100.0, "Groceries")
	expectedTransaction.ID = transactionID

	mockRepo.On("FindByID", transactionID).Return(expectedTransaction, nil)

	result, err := service.GetTransactionByID(transactionID)

	assert.NoError(t, err)
	assert.Equal(t, expectedTransaction, result)

	mockRepo.AssertExpectations(t)
}

func TestTransactionService_GetTransactionByID_NotFound(t *testing.T) {
	service, mockRepo, _, _, _ := setUpTransactionService()

	transactionID := uint(1)

	mockRepo.On("FindByID", transactionID).Return((*models.Transaction)(nil), assert.AnError)

	result, err := service.GetTransactionByID(transactionID)

	assert.Error(t, err)
	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)
}
