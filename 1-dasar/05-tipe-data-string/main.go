package main

import (
	"fmt"
)

func main() {
	var nama_depan string = "Egi"
	var nama_belakang string = "Grivandani"
	nama_lengkap := nama_depan + " " + nama_belakang

	fmt.Println("Nama Lengkap : ", nama_lengkap)
	fmt.Println("Panjang :", len(nama_lengkap))
	fmt.Println("huruf pertama :", string(nama_depan[0]))
	fmt.Println("huruf terakhir :", string(nama_depan[len(nama_depan)-1]))
	fmt.Println("huruf :", nama_depan[:1])
}
