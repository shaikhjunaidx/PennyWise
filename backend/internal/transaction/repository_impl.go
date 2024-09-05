package transaction

import (
	"time"

	"github.com/shaikhjunaidx/pennywise-backend/models"
	"gorm.io/gorm"
)

type WeeklySpending struct {
	Week       int     `json:"week"`
	Year       int     `json:"year"`
	TotalSpent float64 `json:"total_spent"`
}

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

func (r *TransactionRepositoryImpl) FindAllByUsername(username string) ([]*TransactionResponse, error) {
	var transactions []*TransactionResponse

	err := r.DB.Table("transactions").
		Select("transactions.id, transactions.user_id, transactions.category_id, categories.name as category_name, transactions.amount, transactions.description, transactions.transaction_date, transactions.created_at, transactions.updated_at").
		Joins("JOIN users ON users.id = transactions.user_id").
		Joins("JOIN categories ON categories.id = transactions.category_id").
		Where("users.username = ?", username).
		Scan(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepositoryImpl) FindAllByUserIDAndCategoryID(userID, categoryID uint) ([]*TransactionResponse, error) {
	var transactions []*TransactionResponse
	err := r.DB.Where("user_id = ? AND category_id = ?", userID, categoryID).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// func (r *TransactionRepositoryImpl) GetWeeklySpending(userID uint) ([]WeeklySpending, error) {
// 	var weeklySpending []WeeklySpending

// 	err := r.DB.Table("transactions").
// 		Select("YEAR(transaction_date) as year, WEEK(transaction_date, 1) as week, SUM(amount) as total_spent").
// 		Where("user_id = ?", userID).
// 		Group("YEAR(transaction_date), WEEK(transaction_date, 1)").
// 		Order("YEAR(transaction_date) DESC, WEEK(transaction_date, 1) DESC").
// 		Scan(&weeklySpending).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return weeklySpending, nil
// }

// returns last 6 weeks transaction totals
func (r *TransactionRepositoryImpl) GetWeeklySpending(userID uint) ([]WeeklySpending, error) {
	var weeklySpending []WeeklySpending

	cutoffDate := time.Now().AddDate(0, 0, -6*7)

	err := r.DB.Table("transactions").
		Select("YEAR(transaction_date) as year, WEEK(transaction_date, 1) as week, SUM(amount) as total_spent").
		Where("user_id = ? AND transaction_date >= ?", userID, cutoffDate).
		Group("YEAR(transaction_date), WEEK(transaction_date, 1)").
		Order("YEAR(transaction_date) DESC, WEEK(transaction_date, 1) DESC").
		Scan(&weeklySpending).Error

	if err != nil {
		return nil, err
	}

	return weeklySpending, nil
}
