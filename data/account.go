package data

import "time"

type Account struct {
	ID int
	Name string
	AccountType string
	OnBudget bool
}


type AccountBalanceHistory []AccountBalance


type AccountBalance struct {
	ID int
	Account *Account
	Date time.Time
	Balance int
}




