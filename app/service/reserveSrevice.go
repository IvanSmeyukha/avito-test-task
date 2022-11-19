package service

import (
	"avito-test-task/app/models"
	"avito-test-task/app/repository"
	"errors"
	"fmt"
	"time"
)

type ReserveService struct {
	reserveRepo        repository.Reserve
	accountService     Account
	userService        User
	transactionService Transaction
	transactionManager repository.TransactionMng
}

func NewReserveService(reserveRepo repository.Reserve, accountService Account, userService User, transactionService Transaction, transactionManager repository.TransactionMng) *ReserveService {
	return &ReserveService{reserveRepo: reserveRepo, accountService: accountService, userService: userService, transactionService: transactionService, transactionManager: transactionManager}
}

func (r *ReserveService) ReserveMoney(transaction models.Transaction) error {
	r.transactionManager.Begin()

	defer r.transactionManager.Rollback()

	balance, err := r.userService.GetBalance(transaction.UserId)
	if err != nil {
		return err
	}
	if balance < transaction.Value {
		return errors.New("not enough funds on account")
	}

	reserveId, err := r.reserveRepo.AddTransaction(transaction.UserId, transaction.Value)

	if err != nil {
		return err
	}

	date := time.Now()
	datetime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second())

	transaction.ReserveId = reserveId
	transaction.Date = datetime
	transaction.Type = "Reservation"

	err = r.transactionService.AddTransaction(transaction)
	if err != nil {
		return err
	}

	r.transactionManager.Commit()

	return err
}

func (r *ReserveService) WriteOffRevenue(transaction models.Transaction) error {
	r.transactionManager.Begin()

	defer r.transactionManager.Rollback()

	transaction, err := r.transactionService.FindTransaction(transaction)
	if err != nil {
		return err
	}

	err = r.reserveRepo.DeleteTransaction(transaction.ReserveId)
	if err != nil {
		return err
	}

	_, err = r.accountService.AddTransaction(transaction.UserId, transaction.Value*(-1))
	if err != nil {
		return err
	}

	date := time.Now()
	datetime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second())

	transaction.Date = datetime
	transaction.Type = "Write off revenue"

	err = r.transactionService.AddTransaction(transaction)
	if err != nil {
		return err
	}

	r.transactionManager.Commit()

	return nil
}
