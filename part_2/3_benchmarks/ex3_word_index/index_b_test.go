package main

import (
	"testing"
	"math/rand"
	"fmt"
	"strings"
)

func BenchmarkIndex(b *testing.B) {
	if testing.Short() {
        b.Skip("skipping Index benchmark")
    }
	for _, length := range []int{10, 100, 1000, 10000} {
		rand.Seed(0)
		phrase := randomPhrase(length)
		w := MakeWords(phrase)
		name := fmt.Sprintf("Index_Fields-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				w.Index("test")
			}
		})
	}
}

func BenchmarkIndexPrepare(b *testing.B) {
	if testing.Short() {
        b.Skip("skipping Index benchmark")
    }
	for _, length := range []int{10, 100, 1000, 10000} {
		rand.Seed(0)
		phrase := randomPhrase(length)
		w := MakeWords(phrase)
		name := fmt.Sprintf("Index_Fields-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				w.IndexPrepare("test")
			}
		})
	}
}

func randomPhrase(n int) string {
	word := "word"
	var tmpList []string
	for i := 0; i < n; i++ {
		tmpList = append(tmpList, word)
	}
	return strings.Join(tmpList, " ")
}