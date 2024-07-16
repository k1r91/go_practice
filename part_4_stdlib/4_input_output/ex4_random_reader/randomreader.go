package main

import (
	"bufio"
	"fmt"
	"io"
	"crypto/rand"
)

// начало решения

// RandomReader создает читателя, который возвращает случайные байты,
// но не более max штук
type randomReader struct {
	n int
}

func (r *randomReader) Read(p []byte) (int, error) {
	if r.n <= 0 {
		return 0, io.EOF
	}
	if len(p) > r.n {
		p = p[:r.n]
	}
	n, _ := rand.Read(p)
	r.n -= n
	return n, nil
}

func RandomReader(max int) io.Reader {
    return &randomReader{n: max}
}

// конец решения

func main() {
	rnd := RandomReader(5)
	rd := bufio.NewReader(rnd)
	for {
		b, err := rd.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", b)
	}
	fmt.Println()
	// 1 148 253 194 250
	// (значения могут отличаться)
}