package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/k1r911/go_practice/match/glob"
)

func main() {
	pattern, src, err := readInput()
	if err != nil {
		fail(err)
	}
	isMath, err := glob.Match(pattern, src)
	if err != nil {
		fail(err)
	}
	if !isMath {
		os.Exit(0)
	}
	fmt.Println(src)
}


func readInput() (pattern, src string, err error) {
	flag.StringVar(&pattern, "p", "", "pattern to match against")
	flag.Parse()
	if pattern == "" {
		return pattern, src, errors.New("missing pattern")
	}
	src = strings.Join(flag.Args(), "")
	if src == "" {
		return pattern, src, errors.New("missing string to match")
	}
	return pattern, src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}