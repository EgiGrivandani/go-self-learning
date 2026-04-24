package main

import "fmt"

func main() {
	//constanta
	const phi = 3.14
	fmt.Println(phi)

	//Latihan :Buat konstanta untuk hari dalam seminggu menggunakan iota (Senin=1 sampai Minggu=7). Cetak semuanya.

	const (
		senin = iota + 1
		selasa
		rabu
		kamis
		jumat
		sabtu
		minggu
	)

	fmt.Println(senin, selasa, rabu, kamis, jumat, sabtu, minggu)
}
