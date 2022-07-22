package main

import "fmt"

func main() {
	// User defined types

	// Defined the struct
	type profile struct {
		email     string
		firstName string
		lastName  string
	}

	// Declare the struct
	var p profile

	// Initialize the struct
	p = profile{
		email:     "me@example.com",
		firstName: "hi",
		lastName:  "there",
	}

	// Print out the contents
	fmt.Println(p)
}
