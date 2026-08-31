package account

import (
	"CS354P1/customer"
)

type checkingaccount struct {
	account
}

func CheckingAccount(number int, customer *customer.Customer, balance float64) *checkingaccount {
	return &checkingaccount{number: number, customer: customer, balance: balance}
}

func (a *account) Accrue(c chan float64, rate float64) {
	//send 0 to channel since no interest was accrued
	c <- 0
}
