package budget

import "github.com/shaikhjunaidx/pennywise-backend/models"

type BudgetRepository interface {
    Create(budget *models.Budget) error
    Update(budget *models.Budget) error
    DeleteByID(id uint) error
    FindByID(id uint) (*models.Budget, error)
    FindAllByUserID(userID uint) ([]*models.Budget, error)
    FindByUserIDAndCategoryID(userID uint, categoryID *uint, month string, year int) (*models.Budget, error)
}
