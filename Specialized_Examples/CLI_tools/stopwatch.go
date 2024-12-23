package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Press Enter to start the stopwatch...")
	fmt.Scanln()

	start := time.Now()
	fmt.Println("Stopwatch started. Press Enter to stop...")
	fmt.Scanln()

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %.2f seconds\n", elapsed.Seconds())
}
