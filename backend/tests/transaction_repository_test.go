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

func createUser(t *testing.T, db *gorm.DB) *models.User {
	user := &models.User{
		Username:     "john_doe",
		Email:        "john.doe@example.com",
		PasswordHash: "hashed_password",
	}
	err := db.Create(user).Error
	assert.NoError(t, err)
	return user
}

func createCategory(t *testing.T, db *gorm.DB) *models.Category {
	category := &models.Category{
		Name:        "Groceries",
		Description: "Expenses for groceries",
	}
	err := db.Create(category).Error
	assert.NoError(t, err)
	return category
}

func createTransaction(t *testing.T, repo *transaction.TransactionRepositoryImpl, userID, categoryID uint, amount float64, description string) *models.Transaction {
	transaction := &models.Transaction{
		UserID:          userID,
		CategoryID:      categoryID,
		Amount:          amount,
		Description:     description,
		TransactionDate: time.Now(),
	}
	err := repo.Create(transaction)
	assert.NoError(t, err)
	return transaction
}

func TestTransactionRepository_Create(t *testing.T) {
	repo, db := setupTransactionTestRepo(t)
	user := createUser(t, db)
	category := createCategory(t, db)
	transaction := createTransaction(t, repo, user.ID, category.ID, 100.0, "Groceries")

	assert.NotZero(t, transaction.ID)
}

func TestTransactionRepository_FindByID(t *testing.T) {
	repo, db := setupTransactionTestRepo(t)
	user := createUser(t, db)
	category := createCategory(t, db)
	transaction := createTransaction(t, repo, user.ID, category.ID, 100.0, "Groceries")

	foundTransaction, err := repo.FindByID(transaction.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundTransaction)
	assert.Equal(t, transaction.ID, foundTransaction.ID)
}

func TestTransactionRepository_Update(t *testing.T) {
	repo, db := setupTransactionTestRepo(t)
	user := createUser(t, db)
	category := createCategory(t, db)
	transaction := createTransaction(t, repo, user.ID, category.ID, 100.0, "Groceries")

	// Update the transaction
	transaction.Amount = 200.0
	transaction.Description = "Updated Groceries"
	err := repo.Update(transaction)
	assert.NoError(t, err)

	updatedTransaction, err := repo.FindByID(transaction.ID)
	assert.NoError(t, err)
	assert.Equal(t, 200.0, updatedTransaction.Amount)
	assert.Equal(t, "Updated Groceries", updatedTransaction.Description)
}

func TestTransactionRepository_DeleteByID(t *testing.T) {
	repo, db := setupTransactionTestRepo(t)
	user := createUser(t, db)
	category := createCategory(t, db)
	transaction := createTransaction(t, repo, user.ID, category.ID, 100.0, "Groceries")

	err := repo.DeleteByID(transaction.ID)
	assert.NoError(t, err)

	deletedTransaction, err := repo.FindByID(transaction.ID)
	assert.Error(t, err) // Should return an error since the transaction should be deleted
	assert.Nil(t, deletedTransaction)
}

func TestTransactionRepository_FindAllByUsername(t *testing.T) {
	repo, db := setupTransactionTestRepo(t)
	user := createUser(t, db)
	category := createCategory(t, db)

	createTransaction(t, repo, user.ID, category.ID, 50.0, "Dinner")
	createTransaction(t, repo, user.ID, category.ID, 150.0, "Utilities")

	transactions, err := repo.FindAllByUsername(user.Username)
	assert.NoError(t, err)
	assert.Len(t, transactions, 2) // We expect 2 transactions for the user
}
