package transaction

import (
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{DB: db}
}

func (r *TransactionRepositoryImpl) Create(transaction *models.Transaction) error {
	return r.DB.Create(transaction).Error
}

func (r *TransactionRepositoryImpl) Update(transaction *models.Transaction) error {
	return r.DB.Save(transaction).Error
}

func (r *TransactionRepositoryImpl) DeleteByID(id uint) error {
	return r.DB.Delete(&models.Transaction{}, id).Error
}

func (r *TransactionRepositoryImpl) FindByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction

	if err := r.DB.First(&transaction, id).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepositoryImpl) FindAllByUserID(userID uint) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	if err := r.DB.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}
