package main

import "fmt"

func main() {

	var nilaiSiswa = map[string]int{
		"egi":    80,
		"rajesh": 40,
		"kimmy":  30,
	}

	fmt.Println("List: ", nilaiSiswa)

	nilaiSiswa["dani"] = 66
	nilaiSiswa["egi"] = 99
	fmt.Println("List: ", nilaiSiswa)

	delete(nilaiSiswa, "dani")
	fmt.Println("List: ", nilaiSiswa)

	fmt.Println("panjang list: ", len(nilaiSiswa))

	fmt.Println("")

	//Latihan : Buat map kontak berisi 3 nama teman dan nomor teleponnya. Tampilkan semua kontak menggunakan for range, lalu hapus 1 kontak dan cetak ulang.

	var kontak = map[string]string{
		"egi":    "087738393462",
		"rajesh": "08123456789",
		"kimmy":  "08134567890",
	}

	for nama, nomor := range kontak {
		fmt.Println("Nama :", nama)
		fmt.Println("Nomor :", nomor)
	}

	delete(kontak, "egi")
	fmt.Println(kontak)

}
