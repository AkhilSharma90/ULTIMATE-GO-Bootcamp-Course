package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func generatePassword(length int, includeSpecial bool) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const special = "!@#$%^&*()_+[]{}<>?|"
	charSet := letters
	if includeSpecial {
		charSet += special
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			log.Fatal(err)
		}
		password[i] = charSet[idx.Int64()]
	}
	return string(password)
}

func main() {
	var length int
	var includeSpecial bool

	fmt.Print("Enter password length: ")
	fmt.Scan(&length)
	fmt.Print("Include special characters? (true/false): ")
	fmt.Scan(&includeSpecial)

	fmt.Println("Generated Password:", generatePassword(length, includeSpecial))
}
