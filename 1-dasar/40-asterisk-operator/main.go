package main

import "fmt"

func main() {
	var angka int = 42
	var ptr *int = &angka // *int = tipe pointer ke int

	fmt.Println("=== * Sebagai Tipe ===")
	fmt.Println("Tipe ptr   :", fmt.Sprintf("%T", ptr)) // *int
	fmt.Println("Nilai ptr  :", ptr)                    // alamat memori
	fmt.Println("Nilai angka:", angka)                  // 42

	//LATIHAN : Buat pointer ke string berisi nama kamu. Cetak nilainya dengan *ptr, ubah nilainya via pointer, lalu cetak lagi untuk membuktikan perubahan.
	nama := "Egi Grivandani"
	ptrNama := &nama
	*ptrNama = "Kimberly"
	fmt.Println("Nilai ptr  :", nama)

}
