package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrCanceled error = errors.New("canceled")

// начало решения

func withRateLimit(limit int, fn func()) (handle func() error, cancel func()) {
	canceled := false
	delay := time.Duration(1000 / limit * int(time.Millisecond))
	ticker := time.NewTicker(delay)
	handle = func() error {
		if canceled {
			return ErrCanceled
		}
		<-ticker.C
		go fn()
		return nil
	}

	cancel = func() {
		canceled = true
	}
	return handle, cancel
}

// конец решения

func main() {
	work := func() {
		fmt.Print(".")
		time.Sleep(10 * time.Second)
	}

	handle, cancel := withRateLimit(5, work)
	defer cancel()
	start := time.Now()
	const n = 10
	for i := 0; i < n; i++ {
		handle()
	}
	fmt.Println()
	fmt.Printf("%d queries took %v\n", n, time.Since(start))
}
