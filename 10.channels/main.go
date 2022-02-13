package main

import (
	"fmt"
	"net/http"
	"time"
)

// <-c :
// + blocking call
// + sending data with channels
// c <- 5
// myNumber <- c
// fmt.Println(<-c)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.com",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {

		go checkLink(link, c)
	}

	// range get access to every elements of slice, channels,..
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
