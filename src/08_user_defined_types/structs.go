package main

import "fmt"

type webPage struct {
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

func StructsShowcase() {
	// Instantiation method 1
	w1 := &webPage{url: "http://www.oreilly.com/"}

	// Instantiation Method 2
	w2 := new(webPage)

	// Instantiation Method 3
	w3 := webPage{url:"google.com"}

	fmt.Printf("Webpage 1: %v\n", w1)
	fmt.Printf("Webpage 2: %q\n", w2)
	fmt.Printf("Webpage 3: %#v\n", w3)

	fmt.Printf("\nNumberSet struct slice: %#v\n", numberSet)
}
