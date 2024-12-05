package main

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func main() {
	err := &CustomError{Code: 404, Message: "resource not found"}
	fmt.Println(err)
}

//output - Error 404: resource not found
