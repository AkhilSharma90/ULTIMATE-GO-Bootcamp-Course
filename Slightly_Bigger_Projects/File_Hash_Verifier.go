package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Enter file path:")
	var filePath string
	fmt.Scanln(&filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	hash256 := sha256.New()
	hash512 := sha512.New()

	_, err = io.Copy(hash256, file)
	if err != nil {
		fmt.Printf("Error computing hash: %v\n", err)
		return
	}

	_, err = io.Copy(hash512, file)
	if err != nil {
		fmt.Printf("Error computing hash: %v\n", err)
		return
	}

	fmt.Printf("SHA-256: %x\n", hash256.Sum(nil))
	fmt.Printf("SHA-512: %x\n", hash512.Sum(nil))
}
