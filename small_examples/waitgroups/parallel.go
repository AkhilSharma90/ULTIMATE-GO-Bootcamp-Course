package main

import (
	"fmt"
	"sync"
)

func square(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Square of %d is %d\n", n, n*n)
}
func main() {
	var wg sync.WaitGroup
	numbers := []int{2, 4, 6, 8, 10}
	for _, n := range numbers {
		wg.Add(1)
		go square(n, &wg)
	}
	wg.Wait()
	fmt.Println("All numbers processed")
}
