package main

import (
	"errors"
	"fmt"
	"math"
)

type myErr struct {}

func (m myErr) Error() string {
	return fmt.Sprintf("test %d", 5)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("expect x >= 0")
	}
	return math.Sqrt(x), nil
}

func main() {
	for _, x := range []float64{49, -49} {
		if res, err := sqrt(x); err != nil {
			fmt.Printf("sqrt(%v) failed, %v\n", x, err)
		} else {
			fmt.Printf("sqrt(%v) = %v\n", x, res)
		}
	}
}