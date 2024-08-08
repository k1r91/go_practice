package main

import "fmt"

// начало решения

// Map - карта "ключ-значение".
type Map [K comparable, V any] struct {
	st map[K]V
}

// Set устанавливает значение для ключа.
func (m *Map[K, V]) Set(key K, val V) {
	if len(m.st) == 0 {
		m.st = make(map[K]V)
	}
	m.st[key] = val
}

// Get возвращает значение по ключу.
func (m *Map[K, V]) Get(key K) V {
	return m.st[key]
}

// Keys возвращает срез ключей карты.
// Порядок ключей неважен, и не обязан совпадать
// с порядком значений из метода Values.
func (m *Map[K, V]) Keys() []K {
	keys := []K{}
	for k := range m.st {
		keys = append(keys, k)
	}
	return keys
}

// Values возвращает срез значений карты.
// Порядок значений неважен, и не обязан совпадать
// с порядком ключей из метода Keys.
func (m *Map[K, V]) Values() []V {
	vals := []V{}
	for _, v := range m.st {
		vals = append(vals, v)
	}
	return vals
}

// конец решения

func main() {
	m := Map[string, int]{}
	m.Set("one", 1)
	m.Set("two", 2)

	fmt.Println(m.Get("one")) // 1
	fmt.Println(m.Get("two")) // 2

	fmt.Println(m.Keys())   // [one two]
	fmt.Println(m.Values()) // [1 2]
}
