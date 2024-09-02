package budget

import (
	"time"

	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
	"github.com/shaikhjunaidx/pennywise-backend/models"
)



type BudgetService struct {
	Repo        BudgetRepository
	UserService *user.UserService
}

var _ user.UserSignUpBudgetService = (*BudgetService)(nil)

func NewBudgetService(repo BudgetRepository, userService *user.UserService) *BudgetService {
	return &BudgetService{
		Repo:        repo,
		UserService: userService}
}

func (s *BudgetService) CreateBudget(username string, categoryID *uint, amountLimit float64, month string, year int) (*models.Budget, error) {

	user, err := s.UserService.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	budget := &models.Budget{
		UserID:      user.ID,
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

func (s *BudgetService) GetBudgetsForUser(username string) ([]*models.Budget, error) {
	user, err := s.UserService.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	return s.Repo.FindAllByUserID(user.ID)
}

func (s *BudgetService) GetBudgetForUserAndCategory(username string, categoryID *uint, month string, year int) (*models.Budget, error) {
	user, err := s.UserService.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	return s.Repo.FindByUserIDAndCategoryID(user.ID, categoryID, month, year)
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

func (s *BudgetService) CalculateOverallBudget(username string) (*models.Budget, error) {
	user, err := s.UserService.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()
	currentMonth := currentTime.Format("01") // Month as string with leading zero
	currentYear := currentTime.Year()

	budgets, err := s.Repo.FindAllByUserIDAndMonthYear(user.ID, currentMonth, currentYear)
	if err != nil {
		return nil, err
	}

	overallBudget := &models.Budget{
		UserID:          user.ID,
		AmountLimit:     0,
		SpentAmount:     0,
		RemainingAmount: 0,
		BudgetMonth:     currentMonth,
		BudgetYear:      currentYear,
	}

	for _, budget := range budgets {
		overallBudget.AmountLimit += budget.AmountLimit
		overallBudget.SpentAmount += budget.SpentAmount
		overallBudget.RemainingAmount += budget.RemainingAmount
	}

	return overallBudget, nil
}
