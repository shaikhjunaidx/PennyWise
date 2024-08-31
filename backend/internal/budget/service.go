package budget

import (
	"github.com/shaikhjunaidx/pennywise-backend/models"
)

type BudgetService struct {
	Repo BudgetRepository
}

func NewBudgetService(repo BudgetRepository) *BudgetService {
	return &BudgetService{Repo: repo}
}

func (s *BudgetService) CreateBudget(userID uint, categoryID *uint, amountLimit float64, month string, year int) (*models.Budget, error) {
	budget := &models.Budget{
		UserID:      userID,
		CategoryID:  categoryID,
		AmountLimit: amountLimit,
		BudgetMonth: month,
		BudgetYear:  year,
		SpentAmount: 0,
	}
	budget.RemainingAmount = budget.AmountLimit

	if err := s.Repo.Create(budget); err != nil {
		return nil, err
	}

	return budget, nil
}

func (s *BudgetService) UpdateBudget(budgetID uint, amountLimit float64) (*models.Budget, error) {
	budget, err := s.Repo.FindByID(budgetID)
	if err != nil {
		return nil, err
	}

	budget.AmountLimit = amountLimit

	budget.RemainingAmount = amountLimit - budget.SpentAmount

	if err := s.Repo.Update(budget); err != nil {
		return nil, err
	}

	return budget, nil
}

func (s *BudgetService) DeleteBudget(budgetID uint) error {
	return s.Repo.DeleteByID(budgetID)
}

func (s *BudgetService) GetBudgetByID(budgetID uint) (*models.Budget, error) {
	return s.Repo.FindByID(budgetID)
}

func (s *BudgetService) GetBudgetsForUser(userID uint) ([]*models.Budget, error) {
	return s.Repo.FindAllByUserID(userID)
}

func (s *BudgetService) GetBudgetForUserAndCategory(userID uint, categoryID *uint, month string, year int) (*models.Budget, error) {
	return s.Repo.FindByUserIDAndCategoryID(userID, categoryID, month, year)
}

func (s *BudgetService) AddTransactionToBudget(userID uint, categoryID *uint, transactionAmount float64, month string, year int) (*models.Budget, error) {
	budget, err := s.Repo.FindByUserIDAndCategoryID(userID, categoryID, month, year)
	if err != nil {
		return nil, err
	}

	budget.SpentAmount += transactionAmount
	budget.RemainingAmount = budget.AmountLimit - budget.SpentAmount

	if err := s.Repo.Update(budget); err != nil {
		return nil, err
	}

	return budget, nil
}
