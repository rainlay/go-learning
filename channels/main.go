package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//for {
	//	go checkLink(<-c, c)
	//}

	for l := range c {
		// put sleep here will cause main routine block
		go func(link string) {
			time.Sleep(time.Second * 5)
			// wrong example value reference
			go checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "site might be down...")
		c <- link
		return
	}

	fmt.Println(link, "site is up")
	c <- link
}
