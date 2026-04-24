package main

import "fmt"

func main() {
	buah := []string{"apel", "nanas", "anggur"}
	fmt.Println("Buah:", buah)
	fmt.Println("Panjang:", len(buah))   // jumlah elemen
	fmt.Println("Kapasitas:", cap(buah)) // kapasitas internal'

	angka := make([]int, 3, 5)
	angka[0] = 10
	angka[1] = 20
	angka[2] = 30

	fmt.Println("\nSlice make:", angka)
	fmt.Println("Panjang:", len(angka))
	fmt.Println("Kapasitas:", cap(angka))

	sumber := []int{1, 2, 3, 4, 5}
	tujuan := make([]int, len(sumber))

	copy(tujuan, sumber)
	tujuan[0] = 999
	fmt.Println("sumber:", sumber)
	fmt.Println("tujuan:", tujuan)

}
