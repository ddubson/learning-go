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

// Input takes in a "send-only" channel (`chan<-`)
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
	for word := range <-wordChannel {
		fmt.Printf("%s ", word)
	}
	fmt.Println()

	idChan := make(chan int)
	go makeId(idChan)

	// Read from channels -- when reading from a channel that has been closed, a zero value of channel type is returned.
	// We check if the channel is open, then we can interpret the value returned as a value coming from the channel (not just the zero value)
	if v, ok := <-idChan; ok {
		fmt.Printf("%d\n", v)
	} else {
		fmt.Println("channel closed.")
	}
}
