package mocks

import (
	"github.com/shaikhjunaidx/pennywise-backend/internal/transaction"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) Update(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) DeleteByID(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTransactionRepository) FindByID(id uint) (*models.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) FindAllByUsername(username string) ([]*transaction.TransactionResponse, error) {
	args := m.Called(username)
	return args.Get(0).([]*transaction.TransactionResponse), args.Error(1)
}

func (m *MockTransactionRepository) FindAllByUserIDAndCategoryID(userID, categoryID uint) ([]*transaction.TransactionResponse, error) {
	args := m.Called(userID, categoryID)
	return args.Get(0).([]*transaction.TransactionResponse), args.Error(1)
}
