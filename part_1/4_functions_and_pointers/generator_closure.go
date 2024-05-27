package main

import (
	"fmt"
	"sort"
)

func intSequence() func() int {
	i := 0
	return func() int {
		i ++
		return i
	}
}

func main() {
	next := intSequence()
	next2 := intSequence()
	for i := 0; i < 10; i++ {
		fmt.Println(next(), next2())
	}
	a := []int{1, 2, 4, 8, 16, 32, 64, 128}
	x := 53
	fmt.Println(sort.Search(len(a), func(i int) bool { return a[i] >= x}))
}