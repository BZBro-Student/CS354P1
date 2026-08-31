package account

import "CS354P1/customer"

type savingaccount struct {
	account
	interest float64
}

func SavingAccount(number int, customer *customer.Customer, balance float64) *savingaccount {
	return &savingaccount{number: number, customer: customer, balance: balance}
}

func (s *savingaccount) Accrue(c chan float64, rate float64) {
	s.interest += s.balance * rate
	s.balance += s.balance * rate
	// send interest over to the provided channel
	c <- s.interest
}
