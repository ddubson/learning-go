package main

import "fmt"

// Input takes in a "send-only" channel (`chan<-`)
func emit(c chan<- string) {
	words := []string{"The", "Quick", "Brown", "Fox"}

	for _, word := range words {
		// Send to channel
		c <- word
	}

	close(c)
}

func makeId(idChan chan<- int) {
	id := 0
	for {
		// Write id to channel
		idChan <- id
		id += 1
	}
}

func ChannelsShowcase() {
	// Create an unbuffered channel of strings
	wordChannel := make(chan string)

	// go -> starts a goroutine, running concurrent with main() -- non-blocking
	go emit(wordChannel)

	// The emit goroutine pauses until the loop below reads the written value out of the channel -- 'unbuffered'
	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}
	fmt.Println()

	idChan := make(chan int)
	go makeId(idChan)

	// Read from channels.
	fmt.Printf("%d\n", <-idChan)
	fmt.Printf("%d\n", <-idChan)
	fmt.Printf("%d\n", <-idChan)
}
