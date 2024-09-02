package transaction

import (
	"errors"
	"time"

	"github.com/shaikhjunaidx/pennywise-backend/internal/budget"
	"github.com/shaikhjunaidx/pennywise-backend/internal/category"
	"github.com/shaikhjunaidx/pennywise-backend/internal/constants"
	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
	"github.com/shaikhjunaidx/pennywise-backend/models"
)

type TransactionService struct {
	Repo          TransactionRepository
	UserRepo      user.UserRepository
	CategoryRepo  category.CategoryRepository
	BudgetService *budget.BudgetService
}

func NewTransactionService(repo TransactionRepository, userRepo user.UserRepository,
	categoryRepo category.CategoryRepository, budgetService *budget.BudgetService) *TransactionService {
	return &TransactionService{
		Repo:          repo,
		UserRepo:      userRepo,
		CategoryRepo:  categoryRepo,
		BudgetService: budgetService,
	}
}

func (s *TransactionService) AddTransaction(username string, categoryID uint,
	amount float64, description string, transactionDate time.Time) (*models.Transaction, error) {

	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if categoryID == 0 {
		defaultCategory, err := s.CategoryRepo.FindByName(constants.DefaultCategoryName)
		if err != nil {
			return nil, errors.New("default category not found")
		}
		categoryID = defaultCategory.ID
	}

	transaction := &models.Transaction{
		UserID:          user.ID,
		CategoryID:      categoryID,
		Amount:          amount,
		Description:     description,
		TransactionDate: transactionDate,
	}

	if err := s.Repo.Create(transaction); err != nil {
		return nil, err
	}

	if _, err := s.BudgetService.AddTransactionToBudget(user.ID, &categoryID, amount, transactionDate.Month().String(), transactionDate.Year()); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *TransactionService) UpdateTransaction(id uint, amount float64, categoryID uint, description string, transactionDate time.Time) (*models.Transaction, error) {

	transaction, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	oldAmount := transaction.Amount
	oldCategoryID := transaction.CategoryID

	transaction.Amount = amount
	transaction.CategoryID = categoryID
	transaction.Description = description
	transaction.TransactionDate = transactionDate

	if err := s.Repo.Update(transaction); err != nil {
		return nil, err
	}

	if oldCategoryID != categoryID {
		if _, err := s.BudgetService.AddTransactionToBudget(transaction.UserID, &oldCategoryID, -oldAmount, transactionDate.Month().String(), transactionDate.Year()); err != nil {
			return nil, err
		}
		if _, err := s.BudgetService.AddTransactionToBudget(transaction.UserID, &categoryID, amount, transactionDate.Month().String(), transactionDate.Year()); err != nil {
			return nil, err
		}
	} else {
		if _, err := s.BudgetService.AddTransactionToBudget(transaction.UserID, &categoryID, amount-oldAmount, transactionDate.Month().String(), transactionDate.Year()); err != nil {
			return nil, err
		}
	}

	return transaction, nil
}

func (s *TransactionService) DeleteTransaction(transactionID uint) error {

	transaction, err := s.Repo.FindByID(transactionID)
	if err != nil {
		return err
	}

	if err := s.Repo.DeleteByID(transactionID); err != nil {
		return err
	}

	if _, err := s.BudgetService.AddTransactionToBudget(transaction.UserID, &transaction.CategoryID, -transaction.Amount, transaction.TransactionDate.Month().String(), transaction.TransactionDate.Year()); err != nil {
		return err
	}

	return nil
}

func (s *TransactionService) GetTransactionsForUser(username string) ([]*models.Transaction, error) {
	return s.Repo.FindAllByUsername(username)
}

func (s *TransactionService) GetTransactionByID(id uint) (*models.Transaction, error) {
	transaction, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) GetTransactionsByCategoryID(username string, categoryID uint) ([]*models.Transaction, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	transactions, err := s.Repo.FindAllByUserIDAndCategoryID(user.ID, categoryID)
	if err != nil {
		return nil, err
	}

	return transactions, nil

}
