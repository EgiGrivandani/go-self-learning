package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== strings.Builder ===")
	var builder strings.Builder
	builder.WriteString("Halo")
	builder.WriteString(" ")
	builder.WriteString("Dunia")
	builder.WriteString("!")

	hasil := builder.String()
	fmt.Println("Builder:", hasil)

	//sisa nya baca di doc
}
