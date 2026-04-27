package main

import "fmt"

func main() {
	//LATIHAN : Buat variabel slice nil dan map nil. Coba append() ke slice nil (berhasil), lalu coba isi map nil (panic). Gunakan recover untuk menangkap panic-nya.
	defer func() {
		message := recover()
		if message != nil {
			fmt.Println("Panic gan")
		}
	}()

	var x []int = nil
	var y map[string]int = nil
	fmt.Println(x)
	fmt.Println(y)

	//y["xx"] = 1 <= panic

	fmt.Println(append(x, 1))

}
