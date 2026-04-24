package main

import "fmt"

func main() {
	hasilAkhir(30, 30)
}

func hasilAkhir(ujian int, kkm int) {
	lulus := ujian >= kkm
	if lulus {
		fmt.Println("Lulus")
	} else {
		fmt.Println("tidak Lulus")
	}
}
