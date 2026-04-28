package main

import "fmt"

type siswa struct {
	Nama string
	NIM  string
}

type Buku struct {
	Judul   string
	Penulis string
	Halaman int
}

func main() {
	// new(int) → membuat pointer ke int dengan nilai 0
	ptrAngka := new(int)
	fmt.Println("=== new(int) ===")
	fmt.Println("Pointer:", ptrAngka)  // alamat memori
	fmt.Println("Nilai  :", *ptrAngka) // 0 (zero value int)

	ptrSiswa := new(siswa)
	ptrSiswa.Nama = "Egi"
	ptrSiswa.NIM = "14323122311"
	fmt.Println("Nama siswa :", ptrSiswa.Nama)
	fmt.Println("NIM :", ptrSiswa.NIM)

	fmt.Println()
	//LATIHAN : Buat pointer ke struct Buku (Judul, Penulis, Halaman) menggunakan new(). Isi field-nya, lalu cetak hasilnya.
	ptrBook := new(Buku)
	ptrBook.Judul = "Positif Habbit"
	ptrBook.Penulis = "Jason Huang"
	ptrBook.Halaman = 360
	fmt.Println("Judul :", ptrBook.Judul)
	fmt.Println("Penulis :", ptrBook.Penulis)
	fmt.Println("Halaman :", ptrBook.Halaman)

}
