package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Sleeping for 3 seconds...")
	time.Sleep(3 * time.Second)
	fmt.Println("Woke up!")
}

//Sleeping for 3 seconds...
//Woke up!
