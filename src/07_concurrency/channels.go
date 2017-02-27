package main

import "fmt"

func emit(c chan string) {
	words := []string{"The", "Quick", "Brown", "Fox"}

	for _, word := range words {
		// Send to channel
		c <- word
	}

	close(c)
}

func makeId(idChan chan int) {
	id := 0
	for {
		// Write id to channel
		idChan <- id
		id += 1
	}
}

func ChannelsShowcase() {
	wordChannel := make(chan string)

	// go -> starts a goroutine, running conc with main()
	go emit(wordChannel)

	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}
	fmt.Println()


	idChan := make(chan int)
	go makeId(idChan)

	// Read from channels.
	fmt.Printf("%d\n", <- idChan)
	fmt.Printf("%d\n", <- idChan)
	fmt.Printf("%d\n", <- idChan)
}