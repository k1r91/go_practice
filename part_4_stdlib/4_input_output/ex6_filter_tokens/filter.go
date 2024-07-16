package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// TokenReader начитывает токены из источника
type TokenReader interface {
	// ReadToken считывает очередной токен
	// Если токенов больше нет, возвращает ошибку io.EOF
	ReadToken() (string, error)
}

// TokenWriter записывает токены в приемник
type TokenWriter interface {
	// WriteToken записывает очередной токен
	WriteToken(s string) error
}

// начало решения

// FilterTokens читает все токены из src и записывает в dst тех,
// кто проходит проверку predicate
func FilterTokens(dst TokenWriter, src TokenReader, predicate func(s string) bool) (int, error) {
	n := 0
	for {
		token, err := src.ReadToken()
		if err == io.EOF {
			break
		} else if err != nil {
			return n, err
		}
		if predicate(token) {
			err := dst.WriteToken(token)
			if err != nil {
				return n, err
			}
			n += 1
		}
	}
	return n, nil
}

// конец решения

type wordReader struct {
	s *bufio.Scanner
}

type wordWriter struct {
	words []string
}

func NewWordReader(s string) *wordReader {
	reader := strings.NewReader(s)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	return &wordReader{scanner}
}

func NewWordWriter() *wordWriter {
	words := make([]string, 0)
	return &wordWriter{words}
}

func (r *wordReader) ReadToken() (string, error) {
	ok := r.s.Scan()
	if !ok {
		return "", io.EOF
	}
	return r.s.Text(), nil
}

func (w *wordWriter) WriteToken(s string) (error) {
	w.words = append(w.words, s)
	return nil
}

func(w *wordWriter) Words() []string {
	return w.words
}

func main() {
	// Для проверки придется создать конкретные типы,
	// которые реализуют интерфейсы TokenReader и TokenWriter.

	// Ниже для примера используются NewWordReader и NewWordWriter,
	// но вы можете сделать любые на свое усмотрение.

	r := NewWordReader("go is awesome")
	w := NewWordWriter()
	predicate := func(s string) bool {
		return s != "is"
	}
	n, err := FilterTokens(w, r, predicate)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d tokens: %v\n", n, w.Words())
	// 2 tokens: [go awesome]
}