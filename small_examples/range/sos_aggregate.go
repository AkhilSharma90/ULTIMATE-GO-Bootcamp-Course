package main

import "fmt"

type Order struct {
	ID     int
	Amount float64
}

func main() {
	orders := []Order{
		{1, 120.50},
		{2, 450.75},
		{3, 89.90},
	}
	totalAmount := 0.0
	for _, order := range orders {
		totalAmount += order.Amount
	}
	fmt.Printf("Total Order Amount: $%.2f\n", totalAmount)
}
