package main

import (
	"fmt"
	"math/rand"
)

func process(i int) int {
	return i + rand.Int()
}

// BufferedChannelsShowcase
//
// An example of idiomatic use of buffered channels -- knowing exactly how many goroutines that will be launched and
// will write to the channel ahead of time.
func BufferedChannelsShowcase() {
	// # of goroutines to be launched, but also the size of the channel buffer, each routine contributes 1 result
	const conc = 10

	// results channel where all results from concurrent ops are aggregated into
	results := make(chan int, conc)

	// Launch conc # of goroutines
	for i := 0; i < conc; i++ {
		// Take in an integer -- use the index for simplicity
		go func(i int) {
			// Process the int, and send it to channel. do this once.
			results <- process(i)
		}(i)
	}

	var out []int
	for i := 0; i < conc; i++ {
		// Read from channel as results come in -- they will be out of order
		out = append(out, <-results)
	}

	fmt.Printf("Collected results: %+v\n", out)
}
