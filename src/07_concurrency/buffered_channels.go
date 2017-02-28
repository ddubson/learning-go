package main

import (
	"time"
	"math/rand"
	"fmt"
	"sync/atomic"
)

var (
	running int64 = 0
)

func work() {
	atomic.AddInt64(&running, 1)
	fmt.Print("[%d", running)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Print("]")
	atomic.AddInt64(&running, -1)
}

func semaWorker(sema chan bool) {
	<-sema // receive from channel (blocking)

	work()

	sema <- true // put true onto channel
}

func BufferedChannelsShowcase() {
	// At a max of 10 bools
	sema := make(chan bool, 10)

	for i := 0; i < 1000; i++ {
		go semaWorker(sema)
	}

	for i := 0; i < 10; i++ {
		sema <- true
	}

	time.Sleep(30 * time.Second)
}
