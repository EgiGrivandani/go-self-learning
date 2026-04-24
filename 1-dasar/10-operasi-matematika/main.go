package main

import "fmt"

func main() {
	var a int = 10
	var b int = 3

	fmt.Println("tambah", a+b)
	fmt.Println("kurang", a-b)
	fmt.Println("bagi", a/b) //pembulatan ke bawah
	fmt.Println("modulus", a%b)

	var counter int = 0
	counter++                           // tambah 1
	fmt.Println("counter++ →", counter) // 1

	//LATIHAN : Buat kalkulator sederhana: buat 2 variabel a dan b, lalu cetak hasil penjumlahan, pengurangan, perkalian, pembagian, dan sisa baginya.
	calculator(10, 20, "bagi")

}

func calculator(a int, b int, tipe string) {
	if tipe == "tambah" {
		fmt.Println("tambah", a+b)
	} else if tipe == "kurang" {
		fmt.Println("kurang", a-b)
	} else if tipe == "bagi" {
		fmt.Println("bagi", a/b)
	} else if tipe == "modulus" {
		fmt.Println("modulus", a%b)
	}
}
