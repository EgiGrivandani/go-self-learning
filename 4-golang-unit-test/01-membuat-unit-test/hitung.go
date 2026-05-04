package hitung

import "math"

func Tambah(a, b int) int {
	return a + b
}

func Kurang(a, b int) int {
	return a - b
}

func IsBilPrima(n int) bool {

	if n <= 1 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))

	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
