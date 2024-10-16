package data

import (
	"time"
)

type LedgerEntry struct {
	ID        int
	Date      time.Time
	Account   *Account
	Check     string
	TransType string
	Catagory  *Catagory
	Payee     string
	Memo      string
	Amount    int // cents
	Verified  bool
	RawRecord string
}

type Ledger struct {
	Account *Account
	Entries []LedgerEntry
	Balance int
	BalanceDate time.Time
	LastEntry time.Time
}
