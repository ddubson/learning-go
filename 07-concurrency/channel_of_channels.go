package main

import (
	"fmt"
	"time"
)

func emitChannel(chanChannel chan chan string, done chan bool) {
	wordChannel := make(chan string)
	chanChannel <- wordChannel
	defer close(wordChannel)

	words := []string{"The", "Quick", "Brown", "Fox"}
	i := 0
	t := time.NewTimer(1 * time.Millisecond)

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

		case <-t.C:
			// Run for t time and then stop the goroutine
			return
		}
	}
}

func ChannelOfChannelsShowcase() {
	channelCh := make(chan chan string)
	doneCh := make(chan bool)

	go emitChannel(channelCh, doneCh)

	wordCh := <-channelCh

	for word := range wordCh {
		fmt.Printf("%s ", word)
	}
	println()
}
