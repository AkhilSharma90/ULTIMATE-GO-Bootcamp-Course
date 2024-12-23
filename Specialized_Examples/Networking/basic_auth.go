package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/pass", nil)
	auth := "user:pass"
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", "Basic "+encodedAuth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
