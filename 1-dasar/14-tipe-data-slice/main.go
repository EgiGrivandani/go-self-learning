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

	//Latihan : Buat slice kosong daftarBelanja. Tambahkan 4 item menggunakan append(). Lalu cetak semua item dan jumlahnya menggunakan len().
	var daftarBelanja []string
	daftarBelanja = append(daftarBelanja, "sayur")
	daftarBelanja = append(daftarBelanja, "minyak")
	daftarBelanja = append(daftarBelanja, "telur")
	daftarBelanja = append(daftarBelanja, "beras")
	daftarBelanja = append(daftarBelanja, "sabun")
	fmt.Println("\nDaftar Belanja:", daftarBelanja)
	fmt.Println("Jumlah Belanjaan:", len(daftarBelanja))
	fmt.Println("Kapasitas Daftar Belanja:", cap(daftarBelanja))
}
