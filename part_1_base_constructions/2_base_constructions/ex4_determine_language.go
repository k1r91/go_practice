package main

import (
	"fmt"
)

func main() {
	var code string
	var lang string
	fmt.Scan(&code)

	// определите полное название языка по его коду
	// и запишите его в переменную `lang`
	// ...
	switch code {
	case "en":
		lang = "English"
	case "fr":
		lang = "French"
	case "ru", "rus":
		lang = "Russian"
	default:
		lang = "unknown"
	}

	fmt.Println(lang)
}
