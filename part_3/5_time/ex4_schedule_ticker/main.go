package main

import (
	"fmt"
	"time"
)

// начало решения

func schedule(dur time.Duration, fn func()) func() {
	ticker := time.NewTicker(dur)
	canceled := make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-canceled:
				ticker.Stop()
				return
			case <-ticker.C:
				fn()
			}
		}
	}()
	cancel := func() {
		select {
		case canceled <- struct{}{}:
		default:
			return
		}
	}
	return cancel
}

// конец решения

func main() {
	work := func() {
		at := time.Now()
		fmt.Printf("%s: work done\n", at.Format("15:04:05.000"))
	}

	cancel := schedule(50*time.Millisecond, work)
	cancel()
	// хватит на 5 тиков
	time.Sleep(260 * time.Millisecond)
}