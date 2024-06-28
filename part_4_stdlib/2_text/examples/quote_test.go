package quote

import (
	"fmt"
	"strconv"
	"testing"
)

func TestQuoteExamples(t *testing.T) {
	s := "Базы данных"
	fmt.Println(strconv.Quote(s))
	fmt.Println(strconv.QuoteToASCII(s))
	fmt.Println(strconv.Unquote(strconv.Quote(s)))
}