package main

import (
	"fmt"
	"math/rand"
)

// начало решения

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case out <- randomWord(5):
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит

func isUnique(str string) bool {
	checker := make(map[rune]int)
	for _, char := range str {
		checker[char] += 1
		if checker[char] > 1 {
			return false
		}
	}
	return true
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for val := range in {
			if !isUnique(val) {
				continue
			}
			select {
			case out <- val:
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for val := range in {
			select {
			case out <- val + " -> " + Reverse(val):
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan string) <-chan string {
	out := make(chan string)
	go func (){
		defer close(out)
		for {
		select {
		case out <- <- c1:
		case out <- <- c2:
		case <-cancel:
			return
		}
		}
	}()
	/*var wg sync.WaitGroup
	wg.Add(2)
	for _, ch := range []<-chan string{c1, c2} {
		ch := ch
		go func() {
			for val := range ch {
				out <- val
			}
			defer wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()*/
	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan string, n int) {
	for i := 0; i < n; i++ {
		select {
		case val, ok := <- in:
			if !ok {
				return
			}
			fmt.Println(val)
		case <- cancel:
			return
		}
	}
}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
}
