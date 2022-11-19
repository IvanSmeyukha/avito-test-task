package repository

import (
	"github.com/jmoiron/sqlx"
)

// TransactionManager provides requests transactionality on the service layer
type TransactionManager struct {
	db *sqlx.DB
}

func NewTransactionManager(db *sqlx.DB) *TransactionManager {
	return &TransactionManager{db: db}
}

func (s *TransactionManager) Begin() {
	s.db.Exec("START TRANSACTION;")
}

func (s *TransactionManager) Commit() {
	s.db.Exec("COMMIT;")
}

func (s *TransactionManager) Rollback() {
	s.db.Exec("ROLLBACK;")
}
