package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <amount> <currency>")
		return
	}

	amount, _ := strconv.ParseFloat(os.Args[1], 64)
	currency := os.Args[2]

	apiURL := fmt.Sprintf("https://api.exchangerate-api.com/v4/latest/USD")
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching exchange rates:", err)
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)

	rates := data["rates"].(map[string]interface{})
	if rate, ok := rates[currency].(float64); ok {
		fmt.Printf("%.2f USD = %.2f %s\n", amount, amount*rate, currency)
	} else {
		fmt.Println("Currency not supported")
	}
}
