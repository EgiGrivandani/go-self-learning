package main

import "fmt"

type Rupiah int

func main() {
	var hargaBuku Rupiah = 75000
	var hargaPulpen Rupiah = 5000
	totalHarga := hargaPulpen + hargaBuku

	fmt.Println("Harga Buku", hargaBuku)
	fmt.Println("Harga Pulpen", hargaPulpen)
	fmt.Println("Total Buku", totalHarga)
}
