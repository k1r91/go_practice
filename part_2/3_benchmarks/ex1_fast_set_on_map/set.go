package main

// не удаляйте импорты, они используются при проверке
// реализуйте быстрое множество
type IntSet struct {
	elems map[int]bool
}

func MakeIntSet() IntSet {
    return IntSet{make(map[int]bool)}
}

func (s IntSet) Contains(elem int) bool {
    _, ok := s.elems[elem]
	return ok
}

func (s IntSet) Add(elem int) bool {
	_, ok := s.elems[elem]
	if ok {
		return false
	}
    s.elems[elem] = true
	return true
}