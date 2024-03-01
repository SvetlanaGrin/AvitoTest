package repository

import (
	"fmt"
	"github.com/SvetlanaGrin/AvitoTest"
	"github.com/jmoiron/sqlx"
)

type BalancePostgres struct {
	db *sqlx.DB
}

type Balance interface {
	GetById(userid int) (AvitoTest.UserBalance, error)
	//UpdatePlus(userid int, input AvitoTest.UserBalance) error
	//UpdateMinus(userid int, input AvitoTest.UserBalance) error
}

func NewBalancePostgres(db *sqlx.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (b *BalancePostgres) GetById(userid int) (AvitoTest.UserBalance, error) {
	var balanceId AvitoTest.UserBalance
	query := fmt.Sprintf("SELECT tl.balance FROM %s tl WHERE ul.user_id= $1", userBalanceTable)
	err := b.db.Select(&balanceId, query, userid)

	return balanceId, err
}
