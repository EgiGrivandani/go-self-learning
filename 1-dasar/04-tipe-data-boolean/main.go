package main

import "fmt"

func main() {
	var defaultstatus bool
	fmt.Println("Status Default :", defaultstatus)

	var umur int = 20
	var bolehNyoblos bool = umur > 17

	fmt.Println("Bisa Nyoblos :", bolehNyoblos)

	var hujan bool = true
	if hujan {
		fmt.Println("Membawa Payung")
	} else {
		fmt.Println("Tidak Membawa Payung")
	}

}
