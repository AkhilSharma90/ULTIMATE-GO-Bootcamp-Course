package main

import (
	"fmt"
)

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

func generateNumbers(max int, numbers chan int) {
	for i := 2; i <= max; i++ {
		numbers <- i
	}
	close(numbers)
}

func findPrimes(numbers chan int, primes chan int) {
	for num := range numbers {
		if isPrime(num) {
			primes <- num
		}
	}
	close(primes)
}

func main() {
	numbers := make(chan int)
	primes := make(chan int)

	go generateNumbers(100, numbers)
	go findPrimes(numbers, primes)

	fmt.Println("Prime numbers up to 100:")
	for prime := range primes {
		fmt.Println(prime)
	}
}
