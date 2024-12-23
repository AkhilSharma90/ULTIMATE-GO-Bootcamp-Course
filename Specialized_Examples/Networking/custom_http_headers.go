package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	req.Header.Set("User-Agent", "MyCustomClient/1.0")
	req.Header.Set("Authorization", "Bearer mytoken123")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
