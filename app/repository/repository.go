package repository

import (
	"avito-test-task/app/models"
	"github.com/jmoiron/sqlx"
)

type User interface {
	ExistById(int) (bool, error)
	CreateUser(int) error
}

type Account interface {
	AddTransaction(int, float32) (int, error)
	GetBalance(id int) (float32, error)
}

type Transaction interface {
	AddTransaction(transaction models.Transaction) error
	FindTransaction(transaction models.Transaction) (models.Transaction, error)
}

type TransactionMng interface {
	Begin()
	Commit()
	Rollback()
}

type Reserve interface {
	AddTransaction(int, float32) (int, error)
	DeleteTransaction(int) error
}

type Repository struct {
	Reserve
	User
	Account
	Transaction
	TransactionMng
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:           NewUserRepository(db),
		Account:        NewAccountRepository(db),
		Transaction:    NewTransactionRepository(db),
		Reserve:        NewReserveRepository(db),
		TransactionMng: NewTransactionManager(db),
	}
}
