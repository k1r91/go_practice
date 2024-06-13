package text

import (
    "strconv"
    "strings"
)

// AsDigitString returns the number as a string composed of it's digits,
// separated by dashes. E.g. 42513 → 4-2-5-1-3
func AsDigitString(n int) string {
    digits := AsDigits(n)
    parts := make([]string, len(digits))
    for idx, d := range digits {
        parts[idx] = strconv.Itoa(d)
    }
    return strings.Join(parts, "-")
}

// AsRunes returns the number as a slice of it's digits runes.
func AsRunes(n int) []rune {
    return []rune(strconv.Itoa(n))
}

func AsDigits(n int) []int {
    runes := AsRunes(n)
    count := len(runes)
    zero := int('0')
    digits := make([]int, count)
    for idx, char := range runes {
        digits[idx] = int(char) - zero
    }
    return digits
}