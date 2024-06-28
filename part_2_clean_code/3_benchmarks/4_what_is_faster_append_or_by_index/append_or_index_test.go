package append_or_index
import (
	"testing"
	"fmt"
)

func fnAppend(in []int) []int {
	result := make([]int, len(in))
	for _, elem := range in {
		result = append(result, elem)
	}
	return result
}

func fnAppend2(in []int) []int {
	result := make([]int, len(in))
	result = append(result, in...)
	return result
}

func fnAppendByIndex(in []int) []int {
	result := make([]int, len(in))
	for i, elem := range in {
		result[i] = elem
	}
	return result
}

func fnAppendByCopy(in []int) []int {
	result := make([]int, len(in))
	copy(result, in)
	return result
}


func BenchmarkFnAppend(b *testing.B) {
	for _, length := range []int{10000} {
		slice := createSlice(length)
		name := fmt.Sprintf("Allocate-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				fnAppend(slice)
			}
		})
	}
}


func BenchmarkFnAppendByIndex(b *testing.B) {
	for _, length := range []int{10000} {
		slice := createSlice(length)
		name := fmt.Sprintf("Allocate-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				fnAppendByIndex(slice)
			}
		})
	}
}


func BenchmarkFnAppend2(b *testing.B) {
	for _, length := range []int{10000} {
		slice := createSlice(length)
		name := fmt.Sprintf("Allocate-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				fnAppend2(slice)
			}
		})
	}
}


func BenchmarkFnAppendByCopy(b *testing.B) {
	for _, length := range []int{10000} {
		slice := createSlice(length)
		name := fmt.Sprintf("Allocate-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				fnAppendByCopy(slice)
			}
		})
	}
}

func createSlice(length int) []int {
	result := make([]int, 0)
	for i := 0; i < length; i++ {
		result = append(result, i)
	}
	return result
}