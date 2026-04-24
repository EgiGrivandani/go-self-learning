package main

import "fmt"

func main() {
	var harga int = 50000
	var hargaFloat float64 = float64(harga)
	fmt.Println("Harga :", harga)
	fmt.Println("Harga Float :", hargaFloat)

	var harga2 float64 = 900.30
	var hargaInt int = int(harga2)
	fmt.Println("Harga Float :", harga2)
	fmt.Println("Harga Int :", hargaInt)
}
