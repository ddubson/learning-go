package main

import (
	"fmt"
	"io"
	"net/http"
)

// WebPage is a simple 'struct' data type describing the contents of a WebPage
type WebPage struct {
	url  string
	body []byte
	err  error
}

// Inline struct slice declaration
var numberSet = []struct {
	x      int
	y      int
	result int
}{
	{1, 2, 3},
	{2, 2, 4},
	{3, 3, 6},
}

func (w *WebPage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}

	defer resp.Body.Close()

	w.body, err = io.ReadAll(resp.Body)
	if err != nil {
		w.err = err
	}
}

func (w *WebPage) isOK() bool {
	return w.err == nil
}

func userDefinedTypes() {
	// Instantiation method 1
	w1 := &WebPage{url: "https://www.oreilly.com/"}

	// Instantiation Method 2
	w2 := new(WebPage)

	// Instantiation Method 3
	w3 := WebPage{url: "google.com"}

	fmt.Printf("Webpage 1: %v\n", w1)
	fmt.Printf("Webpage 2: %v\n", w2)
	fmt.Printf("Webpage 3: %#v\n", w3)

	fmt.Printf("\nNumberSet struct slice: %#v\n", numberSet)

	// w := &WebPage{url: "https://www.oreilly.com/"}
	w := new(WebPage)

	w.url = "http://www.oreilly.com"
	w.get()

	if w.isOK() {
		fmt.Printf("URL: %s, Error: %s, Length: %d\n", w.url, w.err, len(w.body))
	} else {
		fmt.Print("Something went wrong\n")
	}
}

func main() {
	userDefinedTypes()
}
