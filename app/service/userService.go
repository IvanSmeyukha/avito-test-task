package service

import (
	"avito-test-task/app/models"
	"avito-test-task/app/repository"
	"errors"
	"fmt"
	"time"
)

type UserService struct {
	userRepo           repository.User
	accountService     Account
	transactionService Transaction
	transactionManager repository.TransactionMng
}

func NewUserService(userRepo repository.User, account Account, transaction Transaction, manager repository.TransactionMng) *UserService {
	return &UserService{userRepo: userRepo, accountService: account, transactionService: transaction, transactionManager: manager}
}

func (u *UserService) AddMoney(userId int, amount float32) error {
	u.transactionManager.Begin()

	defer u.transactionManager.Rollback()

	exist, err := u.userRepo.ExistById(userId)
	if !exist {
		u.userRepo.CreateUser(userId)
	}
	_, err = u.accountService.AddTransaction(userId, amount)
	if err != nil {
		return err
	}

	date := time.Now()
	date.Format("")
	datetime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second())

	transaction := models.Transaction{
		UserId: userId,
		Value:  amount,
		Date:   datetime,
		Type:   "Balance replenishment",
	}
	err = u.transactionService.AddTransaction(transaction)
	if err != nil {
		return err
	}

	u.transactionManager.Commit()

	return err
}

func (u *UserService) GetBalance(id int) (float32, error) {
	exist, err := u.userRepo.ExistById(id)
	if !exist || err != nil {
		return 0, errors.New(fmt.Sprintf("user with id = %v was not found", id))
	}
	return u.accountService.GetBalance(id)
}
