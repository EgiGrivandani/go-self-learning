package main

import "fmt"

func filter(numbers []int, fn func(int) bool) []int {
	var result []int
	for _, v := range numbers {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	sapa := func(nama string) string {
		return "hallo bro " + nama
	}

	fmt.Println(sapa("egi"))

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	genap := filter(numbers, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println(genap)

	//LATIHAN : Buat anonymous function yang menerima 2 angka dan mencetak mana yang lebih besar. Panggil langsung tanpa menyimpan ke variabel.
	func(a, b int) {
		if a > b {
			fmt.Println(a, " lebih besar dari ", b)
		} else if b > a {
			fmt.Println(b, " lebih besar dari ", a)
		} else {
			fmt.Println("Nilai sama")
		}
	}(110, 29)

}
