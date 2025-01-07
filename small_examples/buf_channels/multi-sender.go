package main

import (
	"fmt"
	"sync"
)

func sendData(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- id * 10
}
func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sendData(i, ch, &wg)
	}
	wg.Wait()
	close(ch)
	for value := range ch {
		fmt.Println("Received:", value)
	}
}
