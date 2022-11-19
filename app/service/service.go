package service

import (
	"avito-test-task/app/models"
	"avito-test-task/app/repository"
)

type User interface {
	AddMoney(int, float32) error
	GetBalance(int) (float32, error)
}

type Account interface {
	AddTransaction(id int, value float32) (int, error)
	GetBalance(int) (float32, error)
}

type Transaction interface {
	AddTransaction(transaction models.Transaction) error
	FindTransaction(transaction models.Transaction) (models.Transaction, error)
}

type Reserve interface {
	ReserveMoney(models.Transaction) error
	WriteOffRevenue(transaction models.Transaction) error
}

type Service struct {
	User
	Account
	Transaction
	Reserve
}

func NewService(repos *repository.Repository) *Service {
	var service Service
	service.Transaction = NewTransactionService(repos.Transaction)
	service.Account = NewAccountService(repos.Account)
	service.User = NewUserService(repos.User, service.Account, service.Transaction, repos.TransactionMng)
	service.Reserve = NewReserveService(repos.Reserve, service.Account, service.User, service.Transaction, repos.TransactionMng)
	return &service
}
