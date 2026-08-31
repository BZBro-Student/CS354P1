package bank

import (
	"CS354P1/account"
)

type bank struct {
	accounts map[*account.Account]account.Account
}

func NewBank() *bank {
	return &bank{
		accounts: make(map[*account.Account]account.Account),
	}
}
func (b *bank) Add(account account.Account) {
	b.accounts[&account] = account
}

func (b *bank) Accrue(rate float64) float64 {
	//Channel where interest will be sent
	c := make(chan float64)
	var interest float64
	for _, a := range b.accounts {
		go a.Accrue(c, rate)
	}
	for i := 0; i < len(b.accounts); i++ {
		interest += <-c
	}
	return interest
}

func (b *bank) String() string {
	s := ""
	for _, account := range b.accounts {
		s += account.String() + "\n"
	}
	return s
}
