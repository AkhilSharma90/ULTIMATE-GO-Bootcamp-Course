package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <encode|decode> <text>")
		return
	}

	action := os.Args[1]
	text := os.Args[2]

	switch action {
	case "encode":
		fmt.Println("Encoded:", base64.StdEncoding.EncodeToString([]byte(text)))
	case "decode":
		decoded, err := base64.StdEncoding.DecodeString(text)
		if err != nil {
			fmt.Println("Error decoding:", err)
			return
		}
		fmt.Println("Decoded:", string(decoded))
	default:
		fmt.Println("Invalid action. Use 'encode' or 'decode'.")
	}
}
