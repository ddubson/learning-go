package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reader(ch chan int) {
	timer := time.NewTimer(1 * time.Second)
	for {
		select {
		case i := <-ch:
			fmt.Printf("%d\n", i)
		case <-timer.C:
			ch = nil
		}
	}
}

func writer(ch chan int) {
	timer := time.NewTimer(1 * time.Second)

	for {
		select {
		case ch <- rand.Intn(42):
		case <-timer.C:
			ch = nil

		}
	}
}

func NilChannelsShowcase() {
	ch := make(chan int)
	go reader(ch)
	go writer(ch)

	time.Sleep(5 * time.Second)
}
