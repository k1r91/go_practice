package main

import (
	"context"
	"time"
	"fmt"
)

func execute(ctx context.Context, fn func() int) (int, error) {
    ch := make(chan int, 1)

    go func() {
        ch <- fn()
    }()

    select {
    case res := <-ch:
        return res, nil
    case <-ctx.Done():
        return 0, ctx.Err()
    }
}

func main() {
	work := func() int {
		time.Sleep(100 * time.Millisecond)
		return 42
	}
	timeout := time.Duration(50 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := execute(ctx, work)
	fmt.Println(res, err)
}