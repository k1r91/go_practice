package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...
	var res string
	if len(text) <= width {
		res = string([]byte(text))
	} else {
		res = string([]byte(text)[:width]) + "..."
	}
	fmt.Println(res)
}
