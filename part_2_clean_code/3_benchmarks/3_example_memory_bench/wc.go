package wc

import (
    "regexp"
	"strings"
)

// Counter maps words to their counts
type Counter map[string]int

var splitter *regexp.Regexp = regexp.MustCompile(" ")

// WordCountRegexp counts absolute frequencies of words in a string.
// Uses Regexp.Split() to split the string into words.
func WordCountRegexp(s string) Counter {
    counter := make(Counter)
    for _, word := range splitter.Split(s, -1) {
        word = strings.ToLower(word)
        counter[word]++
    }
    return counter
}

func WordCountFields(s string) Counter {
    counter := make(Counter)
    for _, word := range strings.Fields(s) {
        word = strings.ToLower(word)
        counter[word]++
    }
    return counter
}

func WordCountSplit(s string) Counter {
    counter := make(Counter)
    for _, word := range strings.Split(s, " ") {
        word = strings.ToLower(word)
        counter[word]++
    }
    return counter
}

// WordCountAllocate counts absolute frequencies of words in a string.
// Pre-allocates memory for the counter.
func WordCountAllocate(s string) Counter {
    words := strings.Split(s, " ")
    size := len(words) / 2
    if size > 10000 {
        size = 10000
    }
    counter := make(Counter, size)
    for _, word := range words {
        word = strings.ToLower(word)
        counter[word]++
    }
    return counter
}