package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

func delay(dur time.Duration, fn func()) func() {
	timer := time.NewTimer(dur)
	deny := make(chan struct{}, 1)
	go func() {
		select{
			case <-timer.C:
				fn()
			case <- deny:
				return
		}
	}()
	cancel := func(){
		select {
		case deny <- struct{}{}:
		default:
			return
		}
	}
	return cancel
}

// конец решения

func main() {
	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(100*time.Millisecond, work)

	time.Sleep(10 * time.Millisecond)
	if rand.Float32() < 0.5 {
		cancel()
		cancel()
		fmt.Println("delayed function canceled")
	}
	time.Sleep(100 * time.Millisecond)
	cancel()
}
