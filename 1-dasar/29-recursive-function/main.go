package main

import "fmt"

func main() {
	x := factorial(4)
	fmt.Println(x)

	//LATIHAN : Buat recursive function jumlahSampai(n int) int yang menghitung jumlah 1+2+3+...+n. Contoh: jumlahSampai(5) = 15.
	fmt.Println(jumlahSampai(2000))
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	} // base case
	return n * factorial(n-1) // recursive case
}

func jumlahSampai(n int) int {
	for n <= 1 {
		return n
	}
	return n + jumlahSampai(n-1)
}
