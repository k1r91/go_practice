package main

import (
	"errors"
	"fmt"
	"time"
)


func after(dur time.Duration) chan time.Time {
	// also possible to make channel of size = 1, default branch not needed
	done := make(chan time.Time)
	go func() {
		time.Sleep(dur)
		select{
		case done <- time.Now():
		default:
			return
	}		
	}()
	return done
}

func withTimeout(fn func() int, timeout time.Duration) (int, error) {
    var result int

    done := make(chan struct{})
    go func() {
        result = fn()
        close(done)
    }()

    select {
    case <-done:
        return result, nil
    case <-after(timeout):    // (1)
        return 0, errors.New("timeout")
    }
}


func main() {
	fn_long := func() int{
		time.Sleep(10 * time.Second)
		fmt.Println("Long function success")
		return 10
	}
	fn_quick := func() int{
		time.Sleep(10 * time.Millisecond)
		return 10
	}
	_, err := withTimeout(fn_long, 5*time.Millisecond)
	if err != nil {
		fmt.Println("fn long timed out", err)
	} else {
		fmt.Println("fn long success")
	}
	_, err = withTimeout(fn_quick, 50*time.Millisecond)
	if err != nil {
		fmt.Println("fn quick timed out")
	} else {
		fmt.Println("fn quick success")
	}
}