package main

import (
	"strings"
	"testing"
	"go.uber.org/goleak"
)

func helpTestChan(t *testing.T, in <-chan string) {
	for i := 0; i < 100; i++ {
		w := <- in
		parts := strings.Split(w, " -> ")
		if len(parts) != 2 {
			t.Errorf("want 2 parts, got %d parts on word %s, parts are %v", len(parts), w, parts)
		}
		left := parts[0]
		right := parts[1]
		if !isUnique(left) {
			t.Errorf("left word not unique in %s", w)
		}
		if !isUnique(right) {
			t.Errorf("right word not unique in %s", w)
		}
		if Reverse(left) != right {
			t.Errorf("left word is not reversed to right in %s", w)
		}
	}
}

func TestGenerate(t *testing.T) {
	cancel := make(chan struct{})
	defer close(cancel)
	words := generate(cancel)
	for i:=0; i < 5; i ++ {
		word := <- words
		if len(word) != 5 {
			t.Errorf("got '%s' with len %d, want len %d", word, len(word), 5)
		}
	}
	cancel <- struct{}{}
}

func TestIsUnique(t *testing.T) {
	tests := []struct{
		name, in string
		want bool
	}{
		{"corrent unique", "abcdef", true},
		{"not unique", "abbcd", false},
		{"empty string", "", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			got := isUnique(test.in)
			want := test.want
			if got != want {
				t.Errorf("got '%v' on word '%s', want '%v'", got, test.in, want)
			}
		})
	}
}

func TestTakeUnique(t *testing.T) {
	cancel := make(chan struct{})
	defer close(cancel)
	words := generate(cancel)
	uniqueWords := takeUnique(cancel, words)
	for i:=0; i < 100; i++ {
		uniqueWord := <- uniqueWords
		if !isUnique(uniqueWord) {
			t.Errorf("word %s not unique", uniqueWord)
		}
	}
	cancel <- struct{}{}
}

func TestReverse(t *testing.T) {
	tests := []struct{
		name, in, want string
	}{
		{"test #1", "abba", "abba"},
		{"test #2", "abcdef", "fedcba"},
		{"test #3", "", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			got := Reverse(test.in)
			if got != test.want {
				t.Errorf("got %s, want %s", got, test.want)
			}
		})
	}
}


func TestReverseChan(t *testing.T) {
	cancel := make(chan struct{})
	defer close(cancel)
	words := generate(cancel)
	uniqueWords := takeUnique(cancel, words)
	reversedWords := reverse(cancel, uniqueWords)
	helpTestChan(t, reversedWords)
	cancel <- struct{}{}
}

func TestMerge(t *testing.T) {
	defer goleak.VerifyNone(t)
	cancel := make(chan struct{})
	defer close(cancel)
	words := generate(cancel)
	uniqueWords := takeUnique(cancel, words)
	c1 := reverse(cancel, uniqueWords)
	c2 := reverse(cancel, uniqueWords)
	merged := merge(cancel, c1, c2)
	helpTestChan(t, merged)
	cancel <- struct{}{}
}