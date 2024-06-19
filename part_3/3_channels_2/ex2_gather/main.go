package main

import (
	"fmt"
	"time"
)

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	// начало решения
	// выполните все переданные функции,
	// соберите результаты в срез
	// и верните его
	type retval struct {
		idx int
		val any
	}
	result := make([]any, len(funcs))
	done := make(chan retval)
	for i := 0; i < len(funcs); i ++ {
		idx := i
		go func() {
			done <- retval{idx, funcs[idx]()}
		}()
	}
	for i := 0; i < len(funcs); i ++ {
		rv := <- done
		result[rv.idx] = rv.val
	}
	return result
	// конец решения
}

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(3), squared(5), squared(5), squared(6), squared(7), squared(8), squared(9), squared(10), squared(11), squared(12)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}