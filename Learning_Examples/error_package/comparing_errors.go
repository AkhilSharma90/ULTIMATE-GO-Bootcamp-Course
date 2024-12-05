package main

import (
	"errors"
	"fmt"
)

type NetworkError struct {
	Details string
}

func (e *NetworkError) Error() string {
	return e.Details
}

func checkNetwork() error {
	return &NetworkError{Details: "network timeout"}
}

func main() {
	err := checkNetwork()
	var netErr *NetworkError
	if errors.As(err, &netErr) {
		fmt.Println("Network Error Details:", netErr.Details)
	} else {
		fmt.Println("Unknown error occurred")
	}
}

//output - Network Error Details: network timeout
