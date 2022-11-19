package service

import (
	"avito-test-task/app/models"
	"avito-test-task/app/repository"
	"errors"
)

type TransactionService struct {
	transactionRepo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{transactionRepo: repo}
}

func (a *TransactionService) AddTransaction(transaction models.Transaction) error {
	return a.transactionRepo.AddTransaction(transaction)
}

func (a *TransactionService) FindTransaction(transaction models.Transaction) (models.Transaction, error) {
	t, err := a.transactionRepo.FindTransaction(transaction)
	if err != nil {
		return t, errors.New("no reserved funds for his transaction")
	}
	return a.transactionRepo.FindTransaction(transaction)
}
