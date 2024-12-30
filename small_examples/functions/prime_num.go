package main

import "fmt"

// Function to check if a number is prime
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Function to print all prime numbers in a range
func printPrimesInRange(start, end int) {
	fmt.Printf("Prime numbers between %d and %d:\n", start, end)
	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
func main() {
	// Check a specific number
	num := 17
	fmt.Printf("Is %d a prime number? %t\n", num, isPrime(num))
	// Print primes in a range
	printPrimesInRange(10, 50)
}
