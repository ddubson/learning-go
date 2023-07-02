package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPageSize(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return len(body), nil
}

func getter(url string, size chan string) {
	length, err := getPageSize(url)
	if err == nil {
		size <- fmt.Sprintf("%s has length of %d", url, length)
	}
}

func worker(urlCh chan string, sizeCh chan string, id int) {
	for {
		url := <-urlCh
		length, err := getPageSize(url)
		if err == nil {
			sizeCh <- fmt.Sprintf("[worker %d] %s has length %d", id, url, length)
		} else {
			sizeCh <- fmt.Sprintf("[worker %d] Error getting %s: %s", id, url, err)
		}
	}
}

func generator(url string, urlCh chan string) {
	urlCh <- url
}

func MultipleReadersWritersShowcase() {
	urls := []string{"http://www.google.com",
		"http://yahoo.com",
		"http://bing.com",
		"http://bbc.co.uk"}

	sizeCh := make(chan string)

	for _, url := range urls {
		go getter(url, sizeCh)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-sizeCh)
	}

	/// -----

	urlCh := make(chan string)
	sizeCh = make(chan string)

	for i := 0; i < 10; i++ {
		go worker(urlCh, sizeCh, i)
	}

	urlCh <- "http://oreilly.com"

	for _, url := range urls {
		go generator(url, urlCh)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-sizeCh)
	}
}
