package main

// не удаляйте импорты, они используются при проверке
import (
    _"fmt"
    _"math/rand"
    _"os"
    "strings"
    _"testing"
)

type Words struct {
    str string
    words []string
}

func MakeWords(s string) Words {
    words := strings.Fields(s)
    return Words{s, words}
}

func (w Words) Index(word string) int {
    words := strings.Fields(w.str)
    for idx, item := range words {
        if item == word {
            return idx
        }
    }
    return -1
}

func (w Words) IndexPrepare(word string) int {
    for idx, item := range w.words {
        if item == word {
            return idx
        }
    }
    return -1
}


