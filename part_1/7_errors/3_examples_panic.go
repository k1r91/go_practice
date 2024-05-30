package main

import (
	"fmt"
)

func getChar(str string, idx int) byte {
	return str[idx]
}

func protect(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ERROR:", err)
		} else {
			fmt.Println("OK")
		}
	}()
	fn()
}

func work1() {
    defer func() {
        // will recover
        recover()
		fmt.Println("panic oops recovered")
    }()
    panic("oops")
}


func work2() {
    defer func() {
        // will NOT recover
        defer func() {
            recover()
        }()
		panic("oops")
    }()
}


func main() {
	protect(func() {
		c := getChar("test", 10)
		fmt.Println(c)
	})
	protect(func() {
		c := getChar("test", 3)
		fmt.Println(c)
	})
	work1()
	work2()
}