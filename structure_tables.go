package AvitoTest

type CompanyAccounts struct {
	Id     int    `json:"id" db:"id"`
	Number string `json:"number" db:"number_account" binding:"required"  `
	Name   string `json:"name" db:"name_account" `
	Money  int    `json:"money"  db:"money" binding:"required"`
}

type UserBalance struct {
	Id      int `json:"id" db:"id"`
	UserId  int
	Balance int `json:"balance"  db:"balance" binding:"required"`
}
