package repository

import (
	"fmt"
	"github.com/SvetlanaGrin/AvitoTest"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUsers(user AvitoTest.User) (int, error)
	GetUser(username, password string) (AvitoTest.User, error)
}

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUsers(user AvitoTest.User) (int, error) {
	var id int = 1
	query := fmt.Sprintf("INSERT INTO %s (name,username,password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (AvitoTest.User, error) {
	var user AvitoTest.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 and password_hash = $2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
