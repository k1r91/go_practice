package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func toTitleCase(reader *bufio.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	text := scanner.Text()
	var result strings.Builder
	words := strings.Fields(text)
	for i, word := range words {
		lower := strings.ToLower(word)
		rword := []rune(lower)
		result.WriteString(string(unicode.ToUpper(rword[0])))
		result.WriteString(string(rword[1:]))
		if i != len(words) - 1 {
			result.WriteString(" ")
		}
	}
	return result.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(toTitleCase(reader))
}