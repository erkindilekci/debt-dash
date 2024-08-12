package models

type CreditCard struct {
	Id              int
	Name            string
	Limit           int
	CurrentTermDebt float64
	MinimumDebt     float64
}
