package main

import (
	"errors"
	"fmt"
)

func main() {
	nama, panjang := namaLengkap("Budi", "Santoso")
	fmt.Println("Nama   :", nama)
	fmt.Println("Panjang:", panjang, "karakter")

	fmt.Println()

	namaAja, _ := namaLengkap("Andi", "Pratama")
	fmt.Println("Nama saja:", namaAja)

	hasil, err := bagi(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil:", hasil)
	}

	min, max, rata := hitungStatistik([]int{10, 20, 30, 40, 50})
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
	fmt.Println("Rata:", rata)

	//LATIHAN : Buat function konversiSuhu(celsius float64) yang mengembalikan 2 nilai: suhu dalam Fahrenheit dan suhu dalam Kelvin. Rumus: F = C×1.8+32, K = C+273.15.
	c := 27
	f, k := konversiSuhu(float64(c))
	fmt.Println("Celsius :", c)
	fmt.Println("Fahrenheit:", f)
	fmt.Println("Kelvin:", k)
}

func namaLengkap(depan, belakang string) (string, int) {
	lengkap := depan + " " + belakang
	panjang := len(lengkap)
	return lengkap, panjang
}

func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa bagi dengan nol")
	}
	return a / b, nil // nil artinya "tidak ada error"
}

func hitungStatistik(nilai []int) (int, int, float64) {
	min := nilai[0]
	max := nilai[0]
	total := 0

	for _, v := range nilai {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		total += v
	}

	rataRata := float64(total) / float64(len(nilai))
	return min, max, rataRata
}

func konversiSuhu(c float64) (float64, float64) {
	f := (c * 1.8) + 32
	k := c + 273.15
	return f, k
}
