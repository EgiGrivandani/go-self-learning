package main

import "fmt"

func main() {

	// kasus if else statment
	thermal(30)

	//LATIHAN :Buat program yang menampilkan jumlah hari dalam sebuah bulan berdasarkan nama bulan menggunakan switch. Contoh: "Januari" → 31 hari, "Februari" → 29 hari.
	daysInMonth("januari")

}

func thermal(suhu int) {
	switch {
	case suhu < 15:
		fmt.Println("dingin")
	case suhu >= 15 && suhu <= 25:
		fmt.Println("Sejuk")
	case suhu >= 26 && suhu <= 35:
		fmt.Println("Hangat")
	default:
		fmt.Println("panas")
	}
}

func daysInMonth(month string) {
	switch month {
	case "januari", "maret", "mei", "juli", "agustus", "oktober", "desember":
		fmt.Println("31 hari")
	case "april", "juni", "september", "november":
		fmt.Println("30 hari")
	case "februari":
		fmt.Println("29 hari")
	default:
		fmt.Println("bulan tidak valid")
	}
}
