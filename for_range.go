package main

import (
	"fmt"
	"os"
)

func main() {
	// Short initialize both variables to empty strings.
	s, sep := "", ""

	// Create a ranged for loop
	// 'range' produces two values: index, value of the element at index
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
