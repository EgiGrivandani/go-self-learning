package main

import "fmt"

func main() {
	tampilkanKartuNama("Budi", "Programmer", "budi@example.com")
	tampilkanKartuNama("Kimberly", "Desainer Grafis", "kimberly@example.com")

}

func tampilkanKartuNama(nama string, kerjaan string, email string) {
	garisKeras()
	fmt.Println("Nama     : ", nama)
	fmt.Println("Kerjaan  : ", kerjaan)
	fmt.Println("Email    : ", email)
	garisKeras()
}

func garisKeras() {
	fmt.Println("====================")
}
