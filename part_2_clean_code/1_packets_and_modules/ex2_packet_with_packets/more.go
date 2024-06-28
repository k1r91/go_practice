package main

import (
    "fmt"

    "github.com/gothanks/more/numbers"
    "github.com/gothanks/more/text"
)

func main() {
    ok := text.ArePermutation("hello", "lehol")
    fmt.Println("hello / lehol →", ok)

    ok = numbers.IsEven(42)
    fmt.Println("42 is even →", ok)
}