package main

import (
	"fmt"
	"unicode"
)

func checkPasswordStrength(password string) string {
	var hasUpper, hasLower, hasNumber, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	length := len(password)
	if length >= 8 && hasUpper && hasLower && hasNumber && hasSpecial {
		return "Strong"
	} else if length >= 6 && hasUpper && hasLower && (hasNumber || hasSpecial) {
		return "Medium"
	} else {
		return "Weak"
	}
}

func main() {
	fmt.Print("Enter a password: ")
	var password string
	fmt.Scanln(&password)

	strength := checkPasswordStrength(password)
	fmt.Println("Password strength:", strength)
}
