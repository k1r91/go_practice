package examples_test

import (
	"fmt"
	"strconv"
	"testing"
	"strings"
)

func formatList1(items []string) string {
    var str string
    for idx, item := range items {
        str += fmt.Sprintf("%d) %s\n", idx+1, item)
    }
    return str
}

func formatList2(items []string) string {
    strs := make([]string, len(items))
    for idx, item := range items {
        
		strs[idx] = fmt.Sprintf("%d) %s", idx+1, item)
    }
    return strings.Join(strs, "\n")
}

func formatList3(items []string) string {
    var b strings.Builder
    for idx, item := range items {
        b.WriteString(strconv.Itoa(idx + 1))
        b.WriteString(") ")
        b.WriteString(item)
        b.WriteRune('\n')
    }
    return b.String()
}


func BenchmarkFormatList1(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		list := prepareList(length)
		name := fmt.Sprintf("Split-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				formatList1(list)
			}
		})
	}
}

func BenchmarkFormatList2(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		list := prepareList(length)
		name := fmt.Sprintf("Split-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				formatList2(list)
			}
		})
	}
}

func BenchmarkFormatList3(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		list := prepareList(length)
		name := fmt.Sprintf("Split-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				formatList3(list)
			}
		})
	}
}


func prepareList(length int) []string {
	result := make([]string, length)
	for i := 0; i < length; i ++ {
		result[i] = "random phrase"
	}
	return result
}