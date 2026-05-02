package main

import (
	"02-membuat-module/fisika"
	"02-membuat-module/matematika"
	"fmt"
)

func main() {
	fmt.Println(matematika.Tambah(10, 20))
	fmt.Println(fisika.CtoK(23))
}
