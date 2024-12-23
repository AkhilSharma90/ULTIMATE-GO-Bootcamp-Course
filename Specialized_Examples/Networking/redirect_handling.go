package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirected to:", req.URL)
			return nil
		},
	}

	resp, err := client.Get("http://httpbin.org/redirect/3")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Final URL:", resp.Request.URL)
	fmt.Println("Response Status:", resp.Status)
}
