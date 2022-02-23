package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazone.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitUrlChan(url, c)
	}

	for i := 0; i < len(urls); i++ {
		// fmt.Println(<-c)
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

	// for _, url := range urls {
	// 	hitUrl(url)
	// }
}

func hitUrlChan(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}

func hitUrl(url string) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	fmt.Println("url:", url, "status:", status)
}
