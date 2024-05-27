package main

import (
	"fmt"
)

func main() {
	var value any = "hello"
	str := value.(string)
	fmt.Println(str)
	flo, ok := value.(float64)
	fmt.Println(flo, ok)
	switch v := value.(type) {
	case string:
		fmt.Printf("value is string, %#v\n", v)
	case float64:
		fmt.Printf("value is float64, %#v\n", v)
	default:
		fmt.Println("value type is unknown")
	}
}