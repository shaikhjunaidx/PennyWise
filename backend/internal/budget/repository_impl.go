package budget

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/shaikhjunaidx/pennywise-backend/models"
	"gorm.io/gorm"
)

type BudgetRepositoryImpl struct {
	DB *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) *BudgetRepositoryImpl {
	return &BudgetRepositoryImpl{DB: db}
}

func (r *BudgetRepositoryImpl) Create(budget *models.Budget) error {
	return r.DB.Create(budget).Error
}

func (r *BudgetRepositoryImpl) Update(budget *models.Budget) error {
	return r.DB.Save(budget).Error
}

func (r *BudgetRepositoryImpl) DeleteByID(id uint) error {
	return r.DB.Delete(&models.Budget{}, id).Error
}

func (r *BudgetRepositoryImpl) FindByID(id uint) (*models.Budget, error) {
	var budget models.Budget
	if err := r.DB.First(&budget, id).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *BudgetRepositoryImpl) FindAllByUserID(userID uint) ([]*models.Budget, error) {
	var budgets []*models.Budget
	if err := r.DB.Where("user_id = ?", userID).Find(&budgets).Error; err != nil {
		return nil, err
	}
	return budgets, nil
}

func (r *BudgetRepositoryImpl) FindByUserIDAndCategoryID(userID uint, categoryID *uint, month string, year int) (*models.Budget, error) {
	var budget models.Budget
	var monthFormatted string

	if _, err := strconv.Atoi(month); err == nil && len(month) == 2 {
		monthFormatted = month
	} else {
		parsedTime, err := time.Parse("January", month)
		if err != nil {
			return nil, errors.New("invalid month format")
		}
		monthFormatted = fmt.Sprintf("%02d", parsedTime.Month())
	}

	query := r.DB.Where("user_id = ? AND budget_month = ? AND budget_year = ?", userID, monthFormatted, year)

	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	} else {
		query = query.Where("category_id IS NULL")
	}
	if err := query.First(&budget).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *BudgetRepositoryImpl) FindAllByUserIDAndMonthYear(userID uint, month string, year int) ([]*models.Budget, error) {
	var budgets []*models.Budget

	err := r.DB.Where("user_id = ? AND budget_month = ? AND budget_year = ?", userID, month, year).Find(&budgets).Error
	if err != nil {
		return nil, err
	}
	return budgets, nil
}
