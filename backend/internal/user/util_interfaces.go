package user

import "github.com/shaikhjunaidx/pennywise-backend/models"

type UserSignUpCategoryService interface {
	AddCategory(username, name, description string) (*models.Category, error)
	DeleteCategory(username string, id uint) error
}

type UserSignUpBudgetService interface {
	CreateBudget(username string, categoryID *uint, amountLimit float64, month string, year int) (*models.Budget, error)
}
