package main

import (
	_"fmt"
	"regexp"
	"strings"
	"testing"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
func slugify(src string) string {
	src = strings.ToLower(src)
	re := regexp.MustCompile(`[a-z0-9\-]+`)
	words := re.FindAllString(src, -1)
	return strings.Join(words, "-")
}

// конец решения

func Test(t *testing.T) {
	const phrase = "Go Is Awesome!"
	const want = "go-is-awesome"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}