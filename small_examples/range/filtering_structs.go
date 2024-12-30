package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func main() {
	employees := []Employee{
		{"Alice", 55000},
		{"Bob", 48000},
		{"Charlie", 67000},
	}

	var highEarners []Employee
	for _, emp := range employees {
		if emp.Salary > 50000 {
			highEarners = append(highEarners, emp)
		}
	}

	fmt.Println("Employees earning more than $50,000:")
	for _, emp := range highEarners {
		fmt.Printf("%s earns $%.2f\n", emp.Name, emp.Salary)
	}
}
