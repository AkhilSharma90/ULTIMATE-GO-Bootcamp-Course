package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	data := map[string]string{"name": "Alice", "role": "Developer"}
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
