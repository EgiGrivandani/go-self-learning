package main

import "fmt"

func tambahDana(saldo int) func(int) int {
	total := saldo
	return func(tambahan int) int {
		total += tambahan
		return total
	}
}

func buatDiskon(persenDiskon int) func(int) int {
	total := persenDiskon
	return func(harga int) int {
		total = harga - (persenDiskon * harga / 100)
		return total
	}
}

func main() {
	tabungan := tambahDana(100000)
	fmt.Println("Tabung 50rb :", tabungan(50000))
	fmt.Println("Tabung 25rb :", tabungan(25000))
	fmt.Println("Tabung 100rb:", tabungan(100000))

	fmt.Println()

	//LATIHAN : Buat closure buatDiskon(persenDiskon int) yang mengembalikan function untuk menghitung harga setelah diskon. Contoh: diskon20 := buatDiskon(20), lalu diskon20(100000) → 80000.
	diskon20 := buatDiskon(20)
	fmt.Println("Harga 100.000 :", diskon20(100000))
	fmt.Println("Harga 600.000 :", diskon20(600000))
}
