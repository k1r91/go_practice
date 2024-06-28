package examples


import (
	"testing"
	"fmt"
)


func BenchmarkUniqWords(b *testing.B) {
    for _, size := range []int{10, 100, 1000, 10000, 100000} {
        name := fmt.Sprintf("UniqWords-%d", size)
        phrase := randomPhrase(size)
        b.Run(name, func(b *testing.B) {
            for n := 0; n < b.N; n++ {
                UniqWords(phrase)
            }
        })
    }
}

func randomPhrase(size int) string {
	return ""
}