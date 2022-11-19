package repository

import (
	"avito-test-task/app/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) AddTransaction(transaction models.Transaction) error {
	var query string
	if transaction.OrderId == 0 && transaction.ServiceId == 0 {
		query = fmt.Sprintf("INSERT INTO %s (user_id, value, date, type) "+
			"VALUES (%v, %v, \"%s\", \"%s\")",
			transactionsTable,
			transaction.UserId,
			transaction.Value,
			transaction.Date,
			transaction.Type)
	} else if transaction.OrderId != 0 && transaction.ServiceId != 0 {
		query = fmt.Sprintf("INSERT INTO %s (user_id, service_id, order_id, reserve_id, value, date, type) "+
			"VALUES (%v, %v, %v, %v, %v , \"%s\", \"%s\")",
			transactionsTable,
			transaction.UserId,
			transaction.ServiceId,
			transaction.OrderId,
			transaction.ReserveId,
			transaction.Value,
			transaction.Date,
			transaction.Type)
	}
	_, err := r.db.Exec(query)
	return err
}

func (r *TransactionRepository) FindTransaction(transaction models.Transaction) (models.Transaction, error) {
	var query string
	query = fmt.Sprintf("SELECT * FROM %s WHERE service_id = %v AND order_id = %v AND value = %v LIMIT 1",
		transactionsTable,
		transaction.ServiceId,
		transaction.OrderId,
		transaction.Value)
	var res models.Transaction
	err := r.db.QueryRow(query).Scan(
		&res.Id,
		&res.UserId,
		&res.ServiceId,
		&res.OrderId,
		&res.ReserveId,
		&res.Value,
		&res.Date,
		&res.Type)
	return res, err
}
