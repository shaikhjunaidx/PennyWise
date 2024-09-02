package transaction

import "github.com/shaikhjunaidx/pennywise-backend/models"

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	Update(transaction *models.Transaction) error
	DeleteByID(id uint) error
	FindByID(id uint) (*models.Transaction, error)
	FindAllByUsername(username string) ([]*TransactionResponse, error)
	FindAllByUserIDAndCategoryID(userID uint, categoryID uint) ([]*TransactionResponse, error)
}
