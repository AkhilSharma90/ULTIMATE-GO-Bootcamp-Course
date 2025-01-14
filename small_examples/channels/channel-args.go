package main

import "fmt"

func calculateSum(nums []int, ch chan int) {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	ch <- sum
	fmt.Println("Sum is:", sum)
}
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	ch := make(chan int)
	go calculateSum(numbers, ch)

	result := <-ch
	fmt.Println("Sum is:", result)
}
