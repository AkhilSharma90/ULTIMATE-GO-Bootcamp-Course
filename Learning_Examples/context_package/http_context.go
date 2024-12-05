package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func fetchData(ctx context.Context, url string) {
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fetchData(ctx, "https://httpbin.org/delay/3") // Simulates a 3-second delay
}
