package service

import "avito-test-task/app/repository"

type AccountService struct {
	accountRepo repository.Account
}

func NewAccountService(accountRepo repository.Account) *AccountService {
	return &AccountService{accountRepo: accountRepo}
}

func (a *AccountService) AddTransaction(id int, value float32) (int, error) {
	return a.accountRepo.AddTransaction(id, value)
}

func (a *AccountService) GetBalance(id int) (float32, error) {
	return a.accountRepo.GetBalance(id)
}
