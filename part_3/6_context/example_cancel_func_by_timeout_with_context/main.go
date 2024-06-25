package main


import (
	"fmt"
	"time"
	"context"
)

func execute(ctx context.Context, cancel context.CancelFunc, fn func() int, timeoutInSeconds int) (int, error) {
    ch := make(chan int, 1)

    go func() {
        ch <- fn()
    }()
	
	go func() {
		time.Sleep(time.Duration(timeoutInSeconds) * time.Second)
		cancel()
	}()

    select {
    case res := <-ch:
        return res, nil
    case <-ctx.Done():       // (1)
        return 0, ctx.Err()  // (2)
    }
}

func main() {
    // работает в течение 100 мс
    work := func() int {
		time.Sleep(3 * time.Second)
		fmt.Println("work done")
        return 42
    }

    // ждет 50 мс, после этого
    // с вероятностью 50% отменяет работу

    ctx := context.Background()              // (1)
    ctx, cancel := context.WithCancel(ctx)   // (2)
	defer cancel()
    res, err := execute(ctx, cancel, work, 4)           // (5)
    fmt.Println(res, err)
}