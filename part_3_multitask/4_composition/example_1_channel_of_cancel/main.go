package main


import (
	"fmt"
)


func rangeGen(cancel <-chan struct{}, start, stop int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := start; i < stop; i++ {
            select {
            case out <- i:    // (1)
            case <-cancel:    // (2)
                return
            }
        }
    }()
    return out
}


func main() {
    cancel := make(chan struct{})
    defer close(cancel)

    generated := rangeGen(cancel, 41, 46)
    for val := range generated {
        fmt.Println(val)
		// now can interrupt main in any time, gorutine rangeGen will be closed
		if val == 42 {
			cancel <- struct{}{}
		}
    }
}