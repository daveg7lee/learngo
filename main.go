package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://nomadcoders.co/",
	}
	for _, url := range urls {
		err := hitURL(url)
		if err != nil {
			fmt.Println(err)
		}
	}
}

var errRequestFailed = errors.New("Request Failed")

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	fmt.Println(resp.StatusCode)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
