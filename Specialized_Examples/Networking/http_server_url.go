package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	name := keys.Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
