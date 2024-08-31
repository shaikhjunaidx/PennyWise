package test

import (
	"testing"
	"time"

	"github.com/shaikhjunaidx/pennywise-backend/internal/transaction"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/shaikhjunaidx/pennywise-backend/testutils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTransactionTestRepo(t *testing.T) (*transaction.TransactionRepositoryImpl, *gorm.DB) {
	_, tx := testutils.SetupTestDB()
	t.Cleanup(func() {
		tx.Rollback()
	})

	return transaction.NewTransactionRepository(tx), tx
}

func TestTransactionRepository_Create(t *testing.T) {
	repo, _ := setupTransactionTestRepo(t)

	transaction := &models.Transaction{
		UserID:          1,
		CategoryID:      1,
		Amount:          100.0,
		Description:     "Groceries",
		TransactionDate: time.Now(),
	}

	err := repo.Create(transaction)

	assert.NoError(t, err)
	assert.NotZero(t, transaction.ID)
}

func TestTransactionRepository_FindByID(t *testing.T) {
	repo, _ := setupTransactionTestRepo(t)

	transaction := &models.Transaction{
		UserID:          1,
		CategoryID:      1,
		Amount:          100.0,
		Description:     "Groceries",
		TransactionDate: time.Now(),
	}

	err := repo.Create(transaction)
	assert.NoError(t, err)

	foundTransaction, err := repo.FindByID(transaction.ID)
	assert.NoError(t, err)
	assert.Equal(t, transaction.ID, foundTransaction.ID)
	assert.Equal(t, transaction.Description, foundTransaction.Description)
}

func TestTransactionRepository_Update(t *testing.T) {
	repo, _ := setupTransactionTestRepo(t)

	transaction := &models.Transaction{
		UserID:          1,
		CategoryID:      1,
		Amount:          100.0,
		Description:     "Groceries",
		TransactionDate: time.Now(),
	}

	err := repo.Create(transaction)
	assert.NoError(t, err)

	// Update the transaction
	transaction.Amount = 200.0
	transaction.Description = "Updated Groceries"
	err = repo.Update(transaction)
	assert.NoError(t, err)

	updatedTransaction, err := repo.FindByID(transaction.ID)
	assert.NoError(t, err)
	assert.Equal(t, 200.0, updatedTransaction.Amount)
	assert.Equal(t, "Updated Groceries", updatedTransaction.Description)
}

func TestTransactionRepository_DeleteByID(t *testing.T) {
	repo, _ := setupTransactionTestRepo(t)

	transaction := &models.Transaction{
		UserID:          1,
		CategoryID:      1,
		Amount:          100.0,
		Description:     "Groceries",
		TransactionDate: time.Now(),
	}

	err := repo.Create(transaction)
	assert.NoError(t, err)

	err = repo.DeleteByID(transaction.ID)
	assert.NoError(t, err)

	deletedTransaction, err := repo.FindByID(transaction.ID)
	assert.Error(t, err) // Should return an error since the transaction should be deleted
	assert.Nil(t, deletedTransaction)
}

func TestTransactionRepository_FindAllByUserID(t *testing.T) {
	repo, _ := setupTransactionTestRepo(t)

	userID := uint(1)

	transaction1 := &models.Transaction{
		UserID:          userID,
		CategoryID:      1,
		Amount:          50.0,
		Description:     "Dinner",
		TransactionDate: time.Now(),
	}

	transaction2 := &models.Transaction{
		UserID:          userID,
		CategoryID:      2,
		Amount:          150.0,
		Description:     "Utilities",
		TransactionDate: time.Now(),
	}

	err := repo.Create(transaction1)
	assert.NoError(t, err)

	err = repo.Create(transaction2)
	assert.NoError(t, err)

	transactions, err := repo.FindAllByUserID(userID)
	assert.NoError(t, err)
	assert.Len(t, transactions, 2) // We expect 2 transactions for the user
}
