package transaction

import (
	"time"

	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
	"github.com/shaikhjunaidx/pennywise-backend/models"
)

type TransactionService struct {
	Repo     TransactionRepository
	UserRepo user.UserRepository
}

func NewTransactionService(repo TransactionRepository, userRepo user.UserRepository) *TransactionService {
	return &TransactionService{
		Repo: repo,
		UserRepo: userRepo,
	}
}

func (s *TransactionService) AddTransaction(username string, categoryID uint,
	amount float64, description string, transactionDate time.Time) (*models.Transaction, error) {

	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return nil, err
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

	return transaction, nil
}

func (s *TransactionService) UpdateTransaction(id uint, amount float64, categoryID uint, description string, transactionDate time.Time) (*models.Transaction, error) {

	transaction, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	transaction.Amount = amount
	transaction.CategoryID = categoryID
	transaction.Description = description
	transaction.TransactionDate = transactionDate

	if err := s.Repo.Update(transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *TransactionService) DeleteTransaction(transactionID uint) error {

	if err := s.Repo.DeleteByID(transactionID); err != nil {
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
