package account

import (
	"CS354P1/customer"
	"fmt"
)

type Account interface {
	Balance() float64
	Accrue()
	Withdraw()
	toString()
}

type account struct {
	number   int
	customer *customer.Customer
	//No doubles :(
	balance float64
}

func (a *account) Balance() float64 {
	return a.balance
}

func (a *account) Withdraw(amount float64) {
	a.balance -= amount
}

func (a *account) Deposit(amount float64) {
	a.balance += amount
}

func (a *account) String() string {
	return fmt.Sprintf("%d: %s: %f", a.number, a.customer.String(), a.balance)
}
