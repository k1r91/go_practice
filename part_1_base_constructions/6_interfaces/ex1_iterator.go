package main

import (
    "fmt"
)

// element - интерфейс элемента последовательности
// (пустой, потому что элемент может быть любым).
type element interface{}

// iterator - интерфейс, который умеет
// поэлементно перебирать последовательность
type iterator interface {
	val() element
	next() bool
    // определите методы итератора
    // чтобы понять сигнатуры методов - посмотрите,
    // как они используются в функции iterate() ниже
}

// iterate обходит последовательность
// и печатает каждый элемент
func iterate(it iterator) {
    for it.next() {
        curr := it.val()
        fmt.Println(curr)
    }
}

// в этом задании функция main() определена "за кадром",
// не добавляйте ее