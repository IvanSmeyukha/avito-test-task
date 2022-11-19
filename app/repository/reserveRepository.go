package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ReserveRepository struct {
	db *sqlx.DB
}

func NewReserveRepository(db *sqlx.DB) *ReserveRepository {
	return &ReserveRepository{db: db}
}

func (r *ReserveRepository) AddTransaction(userId int, value float32) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, value) VALUES (%v, %v)", reserveTable, userId, value)
	reserveId, err := r.db.Exec(query)
	id, err := reserveId.LastInsertId()
	return int(id), err
}

func (r *ReserveRepository) DeleteTransaction(reserveId int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted = 1 WHERE id = %v", reserveTable, reserveId)
	_, err := r.db.Exec(query)
	return err
}
