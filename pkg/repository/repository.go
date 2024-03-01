package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	Balance
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Balance:       NewBalancePostgres(db),
	}
}
