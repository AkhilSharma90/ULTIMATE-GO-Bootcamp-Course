package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Secure World!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("HTTPS server running on https://localhost:8443")
	http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
}
