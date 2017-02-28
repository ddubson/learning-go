package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

type webPage struct {
	url string
	body []byte
	err error
}

func (w *webPage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}

	defer resp.Body.Close()

	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
	}
}

func (w *webPage) isOK() bool {
	return w.err == nil
}

func main() {
	// w := &webPage{url: "http://www.oreilly.com/"}
	w := new(webPage)

	w.url = "http://www.oreilly.com"
	w.get()

	if w.isOK() {
		fmt.Printf("URL: %s, Error: %s, Length: %d", w.url, w.err, len(w.body))
	} else {
		fmt.Print("Something went wrong")
	}
}
