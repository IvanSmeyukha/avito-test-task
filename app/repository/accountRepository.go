package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (a *AccountRepository) AddTransaction(userId int, value float32) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, value) VALUES (%v, %v)",
		accountsTable,
		userId,
		value)
	accountId, err := a.db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	id, err := accountId.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
	}
	return int(id), err
}

func (a *AccountRepository) GetBalance(userId int) (float32, error) {
	query := fmt.Sprintf(
		"SELECT SUM(value) FROM (SELECT value FROM %s "+
			"WHERE user_id = %v "+
			"UNION ALL SELECT value * (-1) FROM %s "+
			"WHERE user_id = %v AND deleted = 0) AS T",
		accountsTable,
		userId,
		reserveTable,
		userId)
	var balance float32
	err := a.db.Get(&balance, query)
	return balance, err
}
