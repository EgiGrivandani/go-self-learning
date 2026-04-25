package main

import "fmt"

func main() {
	hargaDiskon := hitungDiskon(100000, 20)
	fmt.Println("Harga setelah diskon:", hargaDiskon)

}

func hitungDiskon(harga int, diskonPersen int) int {
	return harga - (harga * diskonPersen / 100)
}
