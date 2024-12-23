package main

import (
	"fmt"
	"net/http"
)

var servers = []string{"http://localhost:8081", "http://localhost:8082"}
var counter int

func loadBalancerHandler(w http.ResponseWriter, r *http.Request) {
	server := servers[counter%len(servers)]
	counter++
	http.Redirect(w, r, server+r.URL.Path, http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/", loadBalancerHandler)
	fmt.Println("Load Balancer running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
