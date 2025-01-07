package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("Request started")
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Request completed")
	case <-ctx.Done():
		fmt.Fprintln(w, "Request canceled")
	}
	fmt.Println("Request ended")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
