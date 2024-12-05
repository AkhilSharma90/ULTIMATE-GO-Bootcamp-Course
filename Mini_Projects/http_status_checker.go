package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// add URL for checking website in bellow.
	links := []string{
		"http://google.com",
		"http://!404!facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	// fmt.Println(<-c) はこのループに入れてもちゃんと走るはず。
	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// fmt.Println(<-c)
	// }

	// for {
	// 	go checkLink(<-c, c)
	// 	// time.Sleep(10 * time.Second)　自作
	// }

	// forの無限ループは見にくいから、以下にリファクタリング
	// function literal, 無名関数を使用する。
	for l := range c {
		go func(link string) {
			time.Sleep(30 * time.Second)
			checkLink(link, c)
		}(l)

	}

}

// Now checkLink func can communicate to channel
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		// c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// c <- "Yes, it's up"
	c <- link
}
