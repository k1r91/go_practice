package main

import (
	"fmt"
	"context"
)

func generate(cxt context.Context, start int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := start; ; i++ {
            select {
            case out <- i:
            case <-cxt.Done():
                return
            }
        }
    }()
    return out
}


func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    generated := generate(ctx, 11)
    for num := range generated {
        fmt.Print(num, " ")
        if num > 14 {
            break
        }
    }
    fmt.Println()
}