package main

import (
	"errors"
	"fmt"
	"net/http"
)

// BufferedChannelsBackpressureShowcase
//
// Backpressure showcase using buffered channels -- limiting the amount of concurrent work that can be done.
func BufferedChannelsBackpressureShowcase() {
	fmt.Println("vvvvvvvvvvv BUF CHAN: BACKPRESSURE vvvvvvvvvvv")

	// Declare a work buffer of 10 goroutines at a time - using backpressure pattern
	pg := New(10)

	http.HandleFunc("/request", func(w http.ResponseWriter, request *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte("Hello!"))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})

	// http.ListenAndServe(":8080", nil)

	fmt.Println("^^^^^^^^^^^ BUF CHAN: BACKPRESSURE ^^^^^^^^^^^^^")
}

type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	// Create an struct channel with buffer size of limit
	ch := make(chan struct{}, limit)

	// Feed struct channel with 'tokens' (indicators of work availability) - 'limit' number of times
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}

	// Return a new pressure gauge
	return &PressureGauge{
		ch: ch,
	}
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	// If there is a work token in the pressuregauge channel, execute 'f'; once done, return the 'token' back into the channel.
	case <-pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("no more capacity")
	}
}
