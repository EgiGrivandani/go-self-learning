package main

import "fmt"

// Sejak Go 1.18, bisa pakai "any" (alias untuk interface{})
func cetakTipe(data any) {
	fmt.Printf("Nilai: %v → Tipe: %T\n", data, data)
}

func deskripsikan(data any) {
	fmt.Printf("Nilai: %v → Tipe: %T\n", data, data)
}

func main() {
	fmt.Println("\n=== Cetak Tipe ===")
	cetakTipe("Golang")
	cetakTipe(2024)

	//LATIHAN : Buat function deskripsikan(data any) yang mencetak tipe dan nilai data. Panggil dengan string, int, float64, bool, dan slice.
	deskripsikan(true)
	deskripsikan("egi")
	deskripsikan(123)
	deskripsikan(123.5)
	deskripsikan([]int{1, 2, 3, 4, 5, 6})
}
