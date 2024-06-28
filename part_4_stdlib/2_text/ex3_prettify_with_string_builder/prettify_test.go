package prettify

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func keys(m map[string]int) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func prettify(m map[string]int) string {
	if len(m) == 0 {
		return "{}"
	} else {
		keys := keys(m)
		if len(m) == 1 {
			return fmt.Sprintf("{ %s: %d }", keys[0], m[keys[0]])
		} else {
			sort.Strings(keys)
			var builder strings.Builder
			builder.WriteString("{\n")
			for _, key := range keys {
				builder.WriteString(fmt.Sprintf("    %s: %d,\n", key, m[key]))
			}
			builder.WriteString("}")
			return builder.String()
		}
	}
}

func TestPrettify(t *testing.T) {
	cases := []struct{
		name string
		in map[string]int
		want string
	}{
		{"1", make(map[string]int), "{}"},
		{"2", map[string]int{"one": 1}, "{ one: 1 }"},
		{"3", map[string]int{"one": 1, "two": 2, "three": 3},
`{
    one: 1,
    three: 3,
    two: 2,
}`,
	},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := prettify(test.in)
			fmt.Println(got)
			fmt.Println(test.want)
			fmt.Println(got == test.want)
			if got != test.want {
				t.Errorf("got %s, want %s", got, test.want)
			}
		})
	}
}