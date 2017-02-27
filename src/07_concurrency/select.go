package main

import (
	"fmt"
	"time"
)

func emitFox(wordChannel chan string, done chan bool) {
	defer close(wordChannel)

	words := []string{"The", "Quick", "Brown", "Fox"}
	i := 0
	t := time.NewTimer(20 * time.Millisecond)

	for {
		select {
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-done:
			done <- true
			return

		case <- t.C: // Run for t time and then stop the goroutine
			return
		}
	}
}

func SelectKeywordShowcase() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emitFox(wordCh, doneCh)

	for word := range wordCh {
		fmt.Printf("%s ", word)
	}
	println()

	//doneCh <- true
}