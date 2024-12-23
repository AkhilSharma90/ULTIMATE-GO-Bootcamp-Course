package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	proxyURL, _ := url.Parse("http://myproxy:8080")
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
	}

	resp, err := client.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
