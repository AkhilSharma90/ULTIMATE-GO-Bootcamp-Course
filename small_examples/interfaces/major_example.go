package main

import (
	"fmt"
)

// PaymentProcessor interface
type PaymentProcessor interface {
	Pay(amount float64) error
}

// Struct for CreditCard
type CreditCard struct {
	CardNumber string
	Balance    float64
}

func (cc *CreditCard) Pay(amount float64) error {
	if cc.Balance < amount {
		return fmt.Errorf("insufficient credit balance on card ending in %s", cc.CardNumber[len(cc.CardNumber)-4:])
	}
	cc.Balance -= amount
	fmt.Printf("Paid %.2f using Credit Card ending in %s. Remaining balance: %.2f\n", amount, cc.CardNumber[len(cc.CardNumber)-4:], cc.Balance)
	return nil
}

// Struct for PayPal
type PayPal struct {
	Email   string
	Balance float64
}

func (pp *PayPal) Pay(amount float64) error {
	if pp.Balance < amount {
		return fmt.Errorf("insufficient PayPal balance for account %s", pp.Email)
	}
	pp.Balance -= amount
	fmt.Printf("Paid %.2f using PayPal account: %s. Remaining balance: %.2f\n", amount, pp.Email, pp.Balance)
	return nil
}

// Struct for Bank Transfer
type BankTransfer struct {
	AccountNumber string
	Balance       float64
}

func (bt *BankTransfer) Pay(amount float64) error {
	if bt.Balance < amount {
		return fmt.Errorf("insufficient balance in bank account ending in %s", bt.AccountNumber[len(bt.AccountNumber)-4:])
	}
	bt.Balance -= amount
	fmt.Printf("Paid %.2f using Bank Transfer from account ending in %s. Remaining balance: %.2f\n", amount, bt.AccountNumber[len(bt.AccountNumber)-4:], bt.Balance)
	return nil
}

// Transaction struct to store transaction history
type Transaction struct {
	Method string
	Amount float64
	Status string
}

// Wallet to manage multiple payment methods and transactions
type Wallet struct {
	PaymentMethods []PaymentProcessor
	Transactions   []Transaction
}

func (w *Wallet) AddPaymentMethod(method PaymentProcessor) {
	w.PaymentMethods = append(w.PaymentMethods, method)
}
func (w *Wallet) Pay(amount float64) {
	for _, method := range w.PaymentMethods {
		err := method.Pay(amount)
		if err == nil {
			w.Transactions = append(w.Transactions, Transaction{
				Method: fmt.Sprintf("%T", method),
				Amount: amount,
				Status: "Success",
			})
			return
		} else {
			fmt.Println("Payment failed:", err)
		}
	}
	w.Transactions = append(w.Transactions, Transaction{
		Method: "All",
		Amount: amount,
		Status: "Failed",
	})
	fmt.Println("All payment methods failed.")
}
func (w *Wallet) ShowTransactions() {
	fmt.Println("\nTransaction History:")
	for i, t := range w.Transactions {
		fmt.Printf("%d. Method: %s, Amount: %.2f, Status: %s\n", i+1, t.Method, t.Amount, t.Status)
	}
}
func main() {
	// Initialize payment methods
	creditCard := &CreditCard{CardNumber: "1234567812345678", Balance: 500}
	payPal := &PayPal{Email: "user@example.com", Balance: 300}
	bankTransfer := &BankTransfer{AccountNumber: "987654321", Balance: 1000}
	// Initialize wallet
	wallet := Wallet{}
	wallet.AddPaymentMethod(creditCard)
	wallet.AddPaymentMethod(payPal)
	wallet.AddPaymentMethod(bankTransfer)
	// Payments
	amounts := []float64{100, 600, 200, 50} // Different amounts to test multiple methods
	for _, amount := range amounts {
		fmt.Printf("\nAttempting to pay %.2f:\n", amount)
		wallet.Pay(amount)
	}
	// Show transaction history
	wallet.ShowTransactions()
}
