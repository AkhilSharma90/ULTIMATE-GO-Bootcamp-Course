package main

import "fmt"

func worker(id int, ch chan string) {
	ch <- fmt.Sprintf("Worker %d finished task", id)
}
func main() {
	ch := make(chan string)
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}
	for i := 1; i <= 3; i++ {
		fmt.Println(<-ch)
	}
}
