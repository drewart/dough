package data

import "time"

type Account struct {
	ID int
	Name string
}


type AccountBalanceHistory []AccountBalance


type AccountBalance struct {
	ID int
	Account *Account
	Date time.Time
	Balance int
}




