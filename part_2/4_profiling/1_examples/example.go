package examples

import (
	"time"
)


func UniqWords(str string) []string {
    words := splitString(str)
    words = sortWords(words)
    words = uniqWords(words)
    return words
}

func splitString(str string) []string {
    time.Sleep(1000 * time.Millisecond)
	return []string{""}
}

func sortWords(words []string) []string {
    time.Sleep(100 * time.Millisecond)
	return []string{""}
}

func uniqWords(words []string) []string {
    time.Sleep(10 * time.Millisecond)
	return []string{""}
}