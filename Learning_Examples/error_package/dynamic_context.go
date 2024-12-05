package main

import (
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero: a=%d, b=%d", a, b)
	}
	return a / b, nil
}

func main() {
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

//output - Error: division by zero: a=10, b=0
