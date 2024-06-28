package main

import (
	"testing"
	"fmt"
)

func Sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func TestMain(m *testing.M) {
    fmt.Println("Testing Sum(n1, n2, ...,nk)...")
	m.Run()
    fmt.Println("Finished testing")
}

func TestSumZero(t *testing.T) {
    if Sum() != 0 {
        t.Errorf("Expected Sum() == 0")
    }
}

func TestSumOne(t *testing.T) {
    if Sum(1) != 1 {
        t.Errorf("Expected Sum(1) == 1")
    }
}

func TestSumPair(t *testing.T) {
    if Sum(1, 2) != 3 {
        t.Errorf("Expected Sum(1, 2) == 3")
    }
    t.Skip()
}

func TestSumMany(t *testing.T) {
    if Sum(1, 2, 3, 4, 5) != 15 {
        t.Errorf("Expected Sum(1, 2, 3, 4, 5) == 15")
    }
}