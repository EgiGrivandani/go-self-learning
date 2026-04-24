package main

import (
	"fmt"
)

func main() {
	var nama string
	fmt.Println("Nama :", nama)

	nama = "Egi"
	fmt.Println("Nama :", nama)

	namaLengkap := "Egi Grivandani"
	fmt.Println("Nama Lengkap :", namaLengkap)

	//multiple variabel
	var alamat, pekerjaan string = "Jakarta", "Programmer"
	fmt.Println("Alamat :", alamat, "Pekerjaan :", pekerjaan)

	var ( //pengelompokan variabel
		nilai  int = 80
		nilai2 int = 78
	)
	fmt.Println("Nilai :", nilai, "Nilai2 :", nilai2)

	//Latihan
	//Buat data diri kamu menggunakan 3 cara berbeda membuat variabel (var dengan tipe, var tanpa tipe, dan :=). Simpan: nama, umur, dan hobi. Cetak semuanya dengan fmt.Println().

	//jawaban
	var namaLengkap2 string = "Egi grivandani" //cara 1: var dengan tipe
	var hobi = "berenang"                      //cara 2: var tanpa tipe
	umur := 25                                 //cara 3: use " := "

	fmt.Println("Nama :", namaLengkap2, "Umur :", umur, "Hobi :", hobi)

}
