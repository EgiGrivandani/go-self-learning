package main

import "fmt"

func main() {
	var data any = []any{1, "xx"}

	switch v := data.(type) {
	case string:
		fmt.Println("String:", v)
	case int:
		fmt.Println("int:", v)
	case bool:
		fmt.Println("bool:", v)
	default:
		fmt.Printf("type -> %T | Data -> %v ", v, v)

	}
}
