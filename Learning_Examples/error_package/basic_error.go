package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("this is a basic error")
	fmt.Println("Error:", err)
}

//output - Error: this is a basic error
