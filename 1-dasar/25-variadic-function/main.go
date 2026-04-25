package main

import "fmt"

func main() {
	nilai := []int{10, 20, 30, 40, 50}
	fmt.Println(jumlahkan(nilai...))
	cetakNilai("Budi", 80, 90, 85)

	//LATIHAN : Buat function cariMax(angka ...int) int yang mengembalikan angka terbesar dari semua parameter yang diberikan.
	numbers := []int{1, 6, 8, 2, 3, 4, 5, 1, 55, 77, 23, 99, 123, 4, 222, 51, 90}
	fmt.Println("List numbers :", numbers)
	fmt.Println("Max number : ", cariMax(numbers...))
}

func jumlahkan(numbers ...int) (total int) {
	for _, num := range numbers {
		total += num
	}
	return

}

func cetakNilai(nama string, nilai ...int) {
	fmt.Println("Nama :", nama)
	fmt.Println("Total Nilai : ", jumlahkan(nilai...))

}

func cariMax(angka ...int) int {
	currentMax := 0
	for _, num := range angka {
		if num > currentMax {
			currentMax = num
		}
	}
	return currentMax
}
