package main

import (
	"CS354P1/account"
	"CS354P1/bank"
	"CS354P1/customer"
	"fmt"
)

func main() {
	ch := make(chan bool)
	bank := bank.NewBank()
	ann := customer.NewCustomer("ann")
	bob := customer.NewCustomer("bob")

	bank.Add(account.CheckingAccount(1, ann, 100.0))
	bank.Add(account.SavingAccount(2, ann, 200.00))
	bank.Add(account.SavingAccount(3, bob, 150.00))

	interest := bank.Accrue(0.02)

	fmt.Printf("%s", bank.String())
	fmt.Printf("Total Interest: %.2f\n", interest)

	bobsChecking := account.CheckingAccount(4, bob, 0)

	for i := 0; i < 1000; i++ {
		go bobsChecking.Deposit(1.00, ch)
	}
	//ensure all threads have completed by checking that each deposit has channeled true
	for i := 0; i < 1000; {
		if <-ch {
			i++
		}
	}

	fmt.Println(bobsChecking.String())

}
