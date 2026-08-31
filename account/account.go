package account

import (
	"CS354P1/customer"
	"fmt"
	"sync"
)

type Account interface {
	Balance() float64
	//Adjusted for goroutine and channels
	Accrue(c chan float64, rate float64)
	Withdraw(amount float64)
	Deposit(amount float64)
	String() string
}

type account struct {
	number   int
	customer *customer.Customer
	balance  float64
	//added a lock for the extra credit
	lock sync.Mutex
}

func (a *account) Balance() float64 {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

/*
	For extra credit: adjusted the function to lock resources during execution
	+= and -= cause a race condition where subroutines are accessing the
	variable at the same time causing read and write operation to overlap
	causing non-determinate behavior

	a.lock.Lock() - locks the resource a once it is available so that only the
	function can modify/read

	a.lock.Unlock() - Allows othere functions to access or lock the resource
	NOTE: Can break the program if Unlock does not happen, causes deadlock error
*/

func (a *account) Withdraw(amount float64) {
	a.lock.Lock()
	a.balance -= amount
	defer a.lock.Unlock()
}

func (a *account) Deposit(amount float64) {
	a.lock.Lock()
	a.balance += amount
	defer a.lock.Unlock()
}

func (a *account) String() string {
	a.lock.Lock()

	defer a.lock.Unlock()
	return fmt.Sprintf("%d: %s: %.2f", a.number, a.customer.String(), a.balance)
}
