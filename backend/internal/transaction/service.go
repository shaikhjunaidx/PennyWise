package transaction

import (
	"time"

	"github.com/shaikhjunaidx/pennywise-backend/models"
)

type TransactionService struct {
	Repo TransactionRepository
}

func NewTransactionService(repo TransactionRepository) *TransactionService {
	return &TransactionService{
		Repo: repo,
	}
}

func (s *TransactionService) AddTransaction(userID, categoryID uint,
	amount float64, description string, transactionDate time.Time) (*models.Transaction, error) {

	transaction := &models.Transaction{
		UserID:          userID,
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

func (s *TransactionService) UpdateTransaction(transactionID uint,
	amount float64, description string) (*models.Transaction, error) {

	transaction, err := s.Repo.FindByID(transactionID)
	if err != nil {
		return nil, err
	}

	transaction.Amount = amount
	transaction.Description = description

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

func (s *TransactionService) GetTransactionForUser(userID uint) ([]*models.Transaction, error) {
	transactions, err := s.Repo.FindAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
