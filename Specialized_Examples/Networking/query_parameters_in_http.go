package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	baseURL, _ := url.Parse("https://example.com/api")
	params := url.Values{}
	params.Add("search", "golang")
	params.Add("limit", "10")
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("URL:", baseURL.String())
	fmt.Println("Response Status:", resp.Status)
}
