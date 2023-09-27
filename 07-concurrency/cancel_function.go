package main

import "fmt"

func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		fmt.Println("closing goroutine")
		close(done)
	}

	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			default:
				ch <- i
			}
		}
		close(ch)
	}()
	return ch, cancel
}

func CancelFunctionShowcase() {
	ch, cancel := countTo(5)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	// goroutine should stop execution
	cancel()
}
