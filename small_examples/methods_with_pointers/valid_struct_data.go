package main

import "fmt"

type Account struct {
	Balance float64
}

// Method to ensure a minimum balance
func (a *Account) EnforceMinimumBalance(min float64) {
	if a.Balance < min {
		a.Balance = min
	}
}
func main() {
	account := Account{Balance: 50}
	account.EnforceMinimumBalance(100)
	fmt.Println("Account balance:", account.Balance) // Output: 100
}
