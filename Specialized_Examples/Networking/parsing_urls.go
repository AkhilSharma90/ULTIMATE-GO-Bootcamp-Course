package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, err := url.Parse("https://example.com/search?q=golang&sort=desc")
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("Host:", u.Host)
	fmt.Println("Path:", u.Path)
	fmt.Println("Query:", u.Query())
}
