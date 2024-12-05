package main

import (
	"errors"
	"fmt"
)

func validate(input int) []error {
	var errs []error
	if input < 0 {
		errs = append(errs, errors.New("input must not be negative"))
	}
	if input > 100 {
		errs = append(errs, errors.New("input exceeds maximum value"))
	}
	return errs
}

func main() {
	errorsList := validate(150)
	if len(errorsList) > 0 {
		fmt.Println("Validation errors:")
		for _, err := range errorsList {
			fmt.Println("-", err)
		}
	}
}

//output
//Validation errors:
//- input must not be negative
//- input exceeds maximum value
