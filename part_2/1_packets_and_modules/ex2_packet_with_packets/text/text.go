package text

import (
    "reflect"
    "sort"
)

// ArePermutation checks if two strings are permutations of each other.
func ArePermutation(first, second string) bool {
    if len(first) != len(second) {
        return false
    }
    runes1st := []rune(first)
    runes2nd := []rune(second)
    sort.Slice(runes1st, func(i, j int) bool { return runes1st[i] < runes1st[j] })
    sort.Slice(runes2nd, func(i, j int) bool { return runes2nd[i] < runes2nd[j] })
    return reflect.DeepEqual(runes1st, runes2nd)
}