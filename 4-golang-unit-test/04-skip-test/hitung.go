package hitung

import "errors"

func Tambah(a, b int) int { return a + b }

func Kurang(a, b int) int { return a - b }

func Kali(a, b int) int { return a * b }

func Bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa membagi dengan nol")
	}
	return a / b, nil
}

func Pangkat(a, n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 0; i < n; i++ {
		hasil *= a
	}
	return hasil
}
