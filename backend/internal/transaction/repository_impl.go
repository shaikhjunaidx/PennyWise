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
	if err := r.DB.Create(transaction).Error; err != nil {
		return err
	}

	if err := r.DB.Preload("User").Preload("Category").
		First(transaction, transaction.ID).Error; err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryImpl) Update(transaction *models.Transaction) error {
	if err := r.DB.Save(transaction).Error; err != nil {
		return err
	}

	if err := r.DB.Preload("User").Preload("Category").
		First(transaction, transaction.ID).Error; err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryImpl) DeleteByID(id uint) error {
	return r.DB.Delete(&models.Transaction{}, id).Error
}

func (r *TransactionRepositoryImpl) FindByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction

	if err := r.DB.Preload("User").Preload("Category").
		First(&transaction, id).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepositoryImpl) FindAllByUsername(username string) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	err := r.DB.Joins("JOIN users ON users.id = transactions.user_id").
		Where("users.username = ?", username).
		Preload("Category").Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepositoryImpl) FindAllByUserIDAndCategoryID(userID, categoryID uint) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.DB.Where("user_id = ? AND category_id = ?", userID, categoryID).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
