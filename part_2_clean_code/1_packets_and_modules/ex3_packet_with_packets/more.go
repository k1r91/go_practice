// -- more/more.go --
package main

import (
    "fmt"

    "github.com/gothanks/morev2/text"
)

func main() {
    digits := text.AsDigits(42513)
    fmt.Println("42513 →", digits)
}