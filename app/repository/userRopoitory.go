package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) ExistById(id int) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE id = %v LIMIT 1)", usersTable, id)
	var exist bool
	err := r.db.Get(&exist, query)
	return exist, err
}

func (r *UserRepository) CreateUser(id int) error {
	query := fmt.Sprintf("INSERT INTO %s (id) VALUES (%v)", usersTable, id)
	_, err := r.db.Exec(query)
	return err
}
