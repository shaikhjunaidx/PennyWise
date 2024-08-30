package transaction

import "github.com/shaikhjunaidx/pennywise-backend/models"

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	Update(transaction *models.Transaction) error
	Delete(transaction *models.Transaction) error
	FindByID(id uint) (*models.Transaction, error)
	FindAllByUserID(userID uint) ([]*models.Transaction, error)
}
