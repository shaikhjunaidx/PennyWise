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

type OverallBudgetResponse struct {
	UserID             uint    `json:"user_id"`
	AmountLimit        float64 `json:"amount_limit"`
	SpentAmount        float64 `json:"spent_amount"`
	RemainingAmount    float64 `json:"remaining_amount"`
	BudgetMonth        string  `json:"budget_month"`
	BudgetYear         int     `json:"budget_year"`
	UncategorizedTotal float64 `json:"uncategorized_total"`
}

type MonthlyBudgetResponse struct {
    Month          string  `json:"month"`
    Year           int     `json:"year"`
    AmountLimit    float64 `json:"amount_limit"`
    SpentAmount    float64 `json:"spent_amount"`
    RemainingAmount float64 `json:"remaining_amount"`
}

type CategoryBudgetHistoryResponse struct {
    CategoryID      uint                    `json:"category_id"`
    History         []MonthlyBudgetResponse `json:"history"`
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

func (s *BudgetService) CalculateOverallBudget(username string) (*OverallBudgetResponse, error) {
	user, err := s.UserService.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()
	currentMonth := currentTime.Format("01")
	currentYear := currentTime.Year()

	budgets, err := s.Repo.FindAllByUserIDAndMonthYear(user.ID, currentMonth, currentYear)
	if err != nil {
		return nil, err
	}

	overallBudget := &OverallBudgetResponse{
		UserID:          user.ID,
		AmountLimit:     0,
		SpentAmount:     0,
		RemainingAmount: 0,
		BudgetMonth:     currentMonth,
		BudgetYear:      currentYear,
	}

	for _, budget := range budgets {
		if budget.AmountLimit > 0 {
			overallBudget.AmountLimit += budget.AmountLimit
			overallBudget.RemainingAmount += budget.RemainingAmount
			overallBudget.SpentAmount += budget.SpentAmount
		} else {
			overallBudget.UncategorizedTotal += budget.SpentAmount
		}
	}

	return overallBudget, nil
}

func (s *BudgetService) GetBudgetHistoryForCategory(username string, categoryID uint) (*CategoryBudgetHistoryResponse, error) {
	user, err := s.UserService.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()

	history := []MonthlyBudgetResponse{}

	for i := 0; i < 4; i++ {
		month := currentTime.AddDate(0, -i, 0).Format("01")
		year := currentTime.AddDate(0, -i, 0).Year()

		budget, err := s.Repo.FindByUserIDAndCategoryID(user.ID, &categoryID, month, year)
		if err != nil {
			// If no budget found, still add it to history with 0 values
			history = append(history, MonthlyBudgetResponse{
				Month:           month,
				Year:            year,
				AmountLimit:     0,
				SpentAmount:     0,
				RemainingAmount: 0,
			})
			continue
		}

		history = append(history, MonthlyBudgetResponse{
			Month:           month,
			Year:            year,
			AmountLimit:     budget.AmountLimit,
			SpentAmount:     budget.SpentAmount,
			RemainingAmount: budget.RemainingAmount,
		})
	}

	return &CategoryBudgetHistoryResponse{
		CategoryID: categoryID,
		History:    history,
	}, nil
}