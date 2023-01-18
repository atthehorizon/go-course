package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
		"https://facebook.com",
	}

	c := make(chan string)
	for _, l := range links {
		go checkStatus(l, c)
	}
	for l := range c {
		go func() {
			time.Sleep(10 * time.Second)
			checkStatus(l, c)
		}()
	}
}

func checkStatus(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "DEAD - HTTP Error:", err)
		c <- l
		return
	}
	fmt.Println(l, "ALIVE")
	c <- l
}
