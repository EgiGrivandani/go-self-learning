package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	x := "ABC@DEFGHIJKLMNOPQRSTUVWXYZ"
	y := strings.NewReader(x)

	data, _ := io.ReadAll(y)
	fmt.Println("ReadAll:", string(data))

	z := strings.NewReader(x)
	limited := io.LimitReader(z, 10)
	lm, _ := io.ReadAll(limited)
	fmt.Println("Limit 10 byte:", string(lm))

	fmt.Println("\n=== Baca per Buffer ===")
	reader2 := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	buf := make([]byte, 5) // buffer 5 byte
	for {
		n, err := reader2.Read(buf)
		if err == io.EOF {
			break // selesai
		}
		fmt.Printf("  Baca %d byte: %s\n", n, string(buf[:n]))
	}

	
}
