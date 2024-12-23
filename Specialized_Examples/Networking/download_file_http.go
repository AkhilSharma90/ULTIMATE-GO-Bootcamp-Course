package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://via.placeholder.com/150")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	file, err := os.Create("downloaded_image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, resp.Body)
}
