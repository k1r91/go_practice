package wc

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkRegexp(b *testing.B) {
    if testing.Short() {
        b.Skip("skipping Regexp benchmark")
    }
	for _, length := range []int{10, 100, 1000, 10000} {
		rand.Seed(0)
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Regexp-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				WordCountRegexp(phrase)
			}
		})
	}
}

func BenchmarkFields(b *testing.B) {
    if testing.Short() {
        b.Skip("skipping Fields benchmark")
    }
	for _, length := range []int{10, 100, 1000, 10000} {
		rand.Seed(0)
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Fields-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				WordCountFields(phrase)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		rand.Seed(0)
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Split-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				WordCountSplit(phrase)
			}
		})
	}
}

func BenchmarkAllocate(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		rand.Seed(0)
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Allocate-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				WordCountAllocate(phrase)
			}
		})
	}
}

// randomPhrase returns a phrase of n random words
func randomPhrase(n int) string {
	word := "word"
	var tmpList []string
	for i := 0; i < n; i++ {
		tmpList = append(tmpList, word)
	}
	return strings.Join(tmpList, " ")
}
