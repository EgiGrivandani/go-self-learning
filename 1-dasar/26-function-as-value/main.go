package main

import "fmt"

func tambah(a, b int) int {
	return a + b
}

func kurang(a, b int) int {
	return a - b
}

func kali(a, b int) int {
	return a * b
}

func bagi(a, b int) int {
	return a / b
}

func toFahrenheit(c float64) float64 {
	return (c * 1.8) + 32
}

func toKelvin(c float64) float64 {
	return c + 273.15
}

func main() {
	var operation func(int, int) int

	operation = tambah
	fmt.Println("Tambah:", operation(10, 5)) // 15

	operation = kurang
	fmt.Println("Kurang:", operation(10, 5)) // 5

	operation = kali
	fmt.Println("Kali  :", operation(10, 5)) // 50

	fmt.Println()

	kalkulator := map[string]func(int, int) int{
		"tambah": tambah,
		"kurang": kurang,
		"kali":   kali,
	}

	fmt.Println("tambah:", kalkulator["tambah"](20, 5))
	fmt.Println("kurang:", kalkulator["kurang"](20, 5))
	fmt.Println("kali  :", kalkulator["kali"](20, 5))

	fmt.Println()

	//LATIHAN : Buat map berisi function untuk konversi suhu: "toFahrenheit" dan "toKelvin". Gunakan map untuk memanggil konversi yang diinginkan.
	calculate := map[string]func(float64) float64{
		"toFahrenheit": toFahrenheit,
		"toKelvin":     toKelvin,
	}

	c := float64(90)
	fmt.Println("c to f:", calculate["toFahrenheit"](c))
	fmt.Println("c to k:", calculate["toKelvin"](c))

}
