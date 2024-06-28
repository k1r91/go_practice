package main

import (
	"fmt"
	"sync"
)

// начало решения

type Counter struct {
	m map[string]int
	mutex *sync.Mutex
}

func (c *Counter) Increment(str string) {
	c.mutex.Lock()
	c.m[str] ++
	c.mutex.Unlock()
}

func (c *Counter) Value(str string) int {
	c.mutex.Lock()
	value := c.m[str]
	c.mutex.Unlock()
	return value
}

func (c *Counter) Range(fn func(key string, val int)) {
	c.mutex.Lock()
	for key, value := range c.m {
		fn(key, value)
	}
	c.mutex.Unlock()
}

func NewCounter() *Counter {
	var mutex sync.Mutex
	return &Counter{make(map[string]int), &mutex}
}

// конец решения

func main() {
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(3)

	increment := func(key string, val int) {
		defer wg.Done()
		for ; val > 0; val-- {
			counter.Increment(key)
		}
	}

	go increment("one", 100)
	go increment("two", 200)
	go increment("three", 300)

	wg.Wait()

	fmt.Println("two:", counter.Value("two"))

	fmt.Print("{ ")
	counter.Range(func(key string, val int) {
		fmt.Printf("%s:%d ", key, val)
	})
	fmt.Println("}")
}
