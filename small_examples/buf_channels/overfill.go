package main

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3 // Uncomment to see deadlock (buffer full and no receiver)
	println(<-ch)
	println(<-ch)
}
