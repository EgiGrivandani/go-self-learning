package main

import "fmt"

func a(name string) string { return "hello " + name }
func c(name string) string { return "selamat siang " + name }
func b(name string, o func(string) string) {
	result := o(name)
	fmt.Println(result)
}

func isGanjil(n int) bool { return n%2 == 1 }
func isGenap(n int) bool  { return n%2 == 0 }

func filter(angka []int, syarat func(int) bool) []int {
	var hasil []int
	for _, v := range angka {
		if syarat(v) {
			hasil = append(hasil, v)
		}
	}
	return hasil
}

func isPass(n int) string {
	result := n >= 75
	if result {
		return "Lulus"
	}
	return "Gagal"
}

func prosesNilai(nilai []int, proses func(int) string) []string {
	var resutl []string
	for _, v := range nilai {
		resutl = append(resutl, proses(v))
	}
	return resutl
}

func main() {
	b("EGI", a)
	b("Kimmy", c)
	fmt.Println()

	angka := []int{1, 2, 3, 4, 5, 6, 3, 2, 1, 45, 75, 4, 3, 5, 33, 4, 6, 7, 2}
	fmt.Println("List angka : ", angka)
	fmt.Println("Result genap: ", filter(angka, isGenap))
	fmt.Println("Result ganjil: ", filter(angka, isGanjil))
	fmt.Println()

	//Latihan : Buat function prosesNilai(nilai []int, proses func(int) string) yang mengubah setiap angka menjadi string "Lulus" (>=75) atau "Gagal" (<75). Kirim function syarat sebagai parameter.
	nilai := []int{45, 78, 50, 77, 75, 90}
	fmt.Println("List nilai : ", nilai)
	fmt.Println("Result: ", prosesNilai(nilai, isPass))
}
