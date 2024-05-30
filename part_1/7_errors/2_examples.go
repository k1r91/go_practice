package main

import (
	"fmt"
	"strings"
)

type lookupError struct {
	src string
	substr string
}

func (e lookupError) Error() string {
	return fmt.Sprintf("'%s' not found in '%s'", e.substr, e.src)
}

func indexOf(src string, substr string) (int, error) {
	idx := strings.Index(src, substr)
	if idx == -1 {
		return -1, lookupError{src, substr}
	}
	return idx, nil
}

func main() {
	src := "some dumb string"
	for _, substr := range []string{"some", " s", "test"} {
		if res, err := indexOf(src, substr); err != nil {
			fmt.Printf("indexOf(%#v, %#v) failed: %v\n", src, substr, err)
		} else {
			fmt.Printf("indexOf(%#v, %#v) = %v\n", src, substr, res)
		}
	}
}