package main

import (
	"fmt"
	"os"
)

func main() {
	// Go only has one loop concept `for` but takes the role of many traditional looping structures

	var counter int = 0;
	for counter < 10 {
		fmt.Print("First Loop\n")
		counter++
	}

	// Traditional single variable loop
	for i := 0; i < 10; i ++ {
		fmt.Print("Second Loop\n")
	}

	// Double variable loop
	for i, j := 0, 1; i < 10; i, j = i + 1, j * 2 {
		fmt.Printf("Third Loop (j=%d)\n", j)
	}

	// Create a ranged for loop
	// 'range' produces two values: index, value of the element at index
	for index, arg := range os.Args[1:] {
		fmt.Printf("%d: %s", index, arg)
	}
}