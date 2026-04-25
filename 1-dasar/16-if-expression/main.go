package main

import "fmt"

func main() {

	//  SHORT STATEMENT
	if panjang := len("egixxxx"); panjang > 3 {
		fmt.Println("Nama panjang")
	} else {
		fmt.Println("Nama pendek")
	}

	//LATIHAN : Buat program yang menerima variabel suhu (int). Cetak: "Dingin" jika < 15, "Sejuk" jika 15-25, "Hangat" jika 26-35, "Panas" jika > 35.
	thermal(45)

}

func thermal(suhu int) {
	if suhu < 15 {
		fmt.Println("dingin")
	} else if suhu >= 15 && suhu <= 25 {
		fmt.Println("Sejuk")
	} else if suhu >= 26 && suhu <= 35 {
		fmt.Println("Hangat")
	} else {
		fmt.Println("Panas")
	}
}
