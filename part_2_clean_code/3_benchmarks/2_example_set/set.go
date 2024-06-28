// -- set.go --
package set

// IntSet implements a mathematical set
// of integers, elements are unique.
type IntSet struct {
    elems *[]int
}

// MakeIntSet creates an empty set.
func MakeIntSet() IntSet {
    elems := []int{}
    return IntSet{&elems}
}

// Contains reports whether an element is within the set.
func (s IntSet) Contains(elem int) bool {
    for _, el := range *s.elems {
        if el == elem {
            return true
        }
    }
    return false
}

// Add adds an element to the set.
func (s IntSet) Add(elem int) bool {
    if s.Contains(elem) {
        return false
    }
    *s.elems = append(*s.elems, elem)
    return true
}