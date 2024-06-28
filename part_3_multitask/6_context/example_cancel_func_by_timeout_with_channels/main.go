package main

import (
	"fmt"
	"time"
	"errors"
)

func execute(cancel chan struct{}, fn func() int, timeoutInSeconds int) (int, error) {
	ch := make(chan int, 1)
	go func() {
		ch <- fn()
	}()
	go func() {
		time.Sleep(time.Duration(timeoutInSeconds) * time.Second)
		close(cancel)
	}()
	select {
	case res := <-ch:
		return res, nil
	case <-cancel:
		return 0, errors.New("timed out")
	}
}

func main() {
	work := func() int {
		time.Sleep(3 * time.Second)
		fmt.Println("work done")
		return 42
	}
	cancel := make(chan struct{})
	res, err := execute(cancel, work, 1)
	fmt.Println(res, err)
}