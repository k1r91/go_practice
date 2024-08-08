package main

import "fmt"

// начало решения

// Avg - накопительное среднее значение.
type Avg[T int|float64] struct {
	counter int
	val T
}

// Add пересчитывает среднее значение с учетом val.
func (a *Avg[T]) Add(val T) *Avg[T] {
	a.counter += 1
	a.val += val
	return a
}

// Val возвращает текущее среднее значение.
func (a *Avg[T]) Val() T {
	if a.counter == 0 {
		return a.val
	}
	return a.val / T(a.counter)
}

// конец решения

func main() {
	intAvg := Avg[int]{}
	intAvg.Add(4).Add(3).Add(2)
	fmt.Println(intAvg.Val()) // 3

	floatAvg := Avg[float64]{}
	floatAvg.Add(4.0).Add(3.0)
	floatAvg.Add(2.0).Add(1.0)
	fmt.Println(floatAvg.Val()) // 2.5
}