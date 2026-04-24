package main

import "fmt"

func main() {
	var buah [3]string
	buah[0] = "Apel"
	buah[1] = "nanas"

	fmt.Println(buah)

	var hewan [4]string = [4]string{"harimau", "gajah", "semut", "koala"}
	fmt.Println(hewan)

	warna := [...]string{"merah", "kuning", "hijau", "biru"}
	fmt.Println(warna)

	//LATIHAN : Buat array nilaiSiswa berisi 5 nilai ujian. Gunakan for range untuk menghitung total dan rata-rata dari semua nilai tersebut.
	var nilaiSiswa [5]int = [5]int{60, 70, 75, 84, 93}
	hitung(nilaiSiswa)
}

func hitung(nilaiSiswa [5]int) {
	total := 0
	for index, value := range nilaiSiswa {
		total += value
		fmt.Println("Index", index, "→ Nilai:", value)
	}
	fmt.Println("Total:", total)

	rataRata := float64(total) / float64(len(nilaiSiswa))
	fmt.Println("Rata:", rataRata)
}
