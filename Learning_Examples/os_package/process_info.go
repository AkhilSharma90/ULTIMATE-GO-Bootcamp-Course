package main

import (
	"fmt"
	"os"
)

func main() {
	// Getting process ID and parent process ID
	fmt.Println("Current Process ID:", os.Getpid())
	fmt.Println("Parent Process ID:", os.Getppid())
}
