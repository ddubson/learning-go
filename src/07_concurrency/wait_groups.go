package main

import (
	"fmt"
	"time"
	"sync"
)

func WaitGroupsShowcase() {
	fmt.Println("----------- WAIT GROUPS -------------")
	queue := make(chan string)
	var w sync.WaitGroup
	w.Add(2)

	go func(queue chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("func 1")
		time.Sleep(2 * time.Second)
		fmt.Println("func 1")
		queue <- "done func 1"
	}(queue, &w)

	go func(queue chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("func 2")
		time.Sleep(1 * time.Second)
		fmt.Println("func 2")
		queue <- "done func 2"
	}(queue, &w)

	go func(queue chan string) {
		for text := range queue {
			fmt.Println(text)
		}
	}(queue)

	// Blocks until all groups are finished executing
	w.Wait()
}
