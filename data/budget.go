package data

import "time"

type Budget struct {
	ID int
	Name string
}

type BudgetCategory struct {
	Budget *Budget
	Catagory *Catagory
	Month time.Time
	Amount int
	Active int
}