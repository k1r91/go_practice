package main

import (
	"fmt"
	"strings"
	"testing"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
/*func slugify(src string) string {
	lower := strings.Trim(strings.ToLower(src), " ")
	fields := strings.Fields(lower)
	result := make([]string, len(fields))
	specialChars := []int{39, 95, 92, 47, 46}
	strings.FieldsFunc()
	for i, word := range fields {
		processedWord := make([]string, len(word))
		for j, r := range word {
			if (r >= 97 && r <= 122) || (r >= 48 && r <= 57) || r == 45 {
				processedWord[j] = string(r)
			} else {
				for _, c := range specialChars {
					if int(r) == c {
						processedWord[j] = "-"
						break
					}
				}
			}
		}
		if len(processedWord) > 1 {
			result[i] = strings.Trim(strings.Join(processedWord, ""), "-")
		} else {
			result[i] = strings.Join(processedWord, "")
		}
	}
	return strings.Trim(strings.Join(result, "-"), "-")
}*/
/*func slugify(src string) string {
    src = strings.ToLower(src)
    words := strings.FieldsFunc(src, func(r rune) bool {
        return (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') && r != '-'
    })
    return strings.Join(words, "-")
}*/

func slugify(src string) string {
    src = strings.ToLower(src)
	length := len(src)
	var b strings.Builder
	var wordLen int
	b.Grow(length)
	//bsrc := []byte(src)
	//fmt.Println(src, length)
	isSafe := func(c byte) bool {
		if (c >= 97 && c <= 122) || c == 45 || (c >= 48 && c <= 57) {
			return true
		}
		return false
	}
	for i := 0; i < length; {
		j := i
		//fmt.Printf("another word, i=%d\n", i)
		wordLen = 0
		for j < length && isSafe(src[j]) {
			b.WriteByte(src[j])
			wordLen ++
			j += 1
		}
		if j < length - 1 && wordLen > 0 {
			b.WriteByte(45)
		}
		i = j + 1
	}
	if wordLen == 0 {
    	return strings.Trim(b.String(), "-")
	} else {
		return b.String()
	}
}

// конец решения
func Test(t *testing.T) {
	fmt.Println([]byte("'"))
	fmt.Println([]rune{'a', 'z', '0', '9', ' ', '-', '\'', '_', '\\', '/', '.'})
	const phrase = "Debugging Go code (a status report)"
	const want = "debugging-go-code-a-status-report"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}

func Test2(t *testing.T) {
	const phrase = "Go-Is-Awesome"
	const want = "go-is-awesome"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}


func Test3(t *testing.T) {
	const phrase = "Go's New Brand"
	const want = "go-s-new-brand"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}


func Test4(t *testing.T) {
	const phrase = "Arrays, slices (and strings): The mechanics of 'append'"
	const want = "arrays-slices-and-strings-the-mechanics-of-append"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}

func Test5(t *testing.T) {
	const phrase = "Go - Is - Awesome"
	const want = "go---is---awesome"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}

func Test6(t *testing.T) {
	const phrase = "Hello, 中国!"
	const want = "hello"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}

func Test7(t *testing.T) {
	const phrase = "Tz6t5bx S9zne Fw-6i Giv0f F894; Tp-"
	const want = "tz6t5bx-s9zne-fw-6i-giv0f-f894-tp-"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}