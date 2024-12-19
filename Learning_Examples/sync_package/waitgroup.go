package main

import (
	"fmt"
	"sync"
)

func printNumber(wg *sync.WaitGroup, number int) {
	defer wg.Done()
	fmt.Println(number)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go printNumber(&wg, i)
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}


//sync.WaitGroup is used to block the main function until all goroutines have completed their work.