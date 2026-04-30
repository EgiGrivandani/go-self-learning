package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "122@3"
	conv, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println("Err : ", err)
	} else {
		fmt.Println("Hasil:", conv, "| Tipe:", fmt.Sprintf("%T", conv))
	}
}
