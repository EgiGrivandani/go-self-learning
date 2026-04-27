package main

import "fmt"

type address struct {
	Kota, Provinsi, Negara string
}

func main() {
	nama := "Egi"
	nama2 := &nama

	fmt.Println(nama)
	fmt.Println(nama2)

	*nama2 = "Budi"
	fmt.Println(nama)

	address1 := address{"Pwk", "Jabar", "Indo"}
	address2 := &address1

	address2.Kota = "Subang"
	fmt.Println(address1)
	fmt.Println(address2)

	//LATIHAN : Buat 2 variabel int, lalu buat pointer untuk masing-masing. Tukar nilainya menggunakan pointer saja (tanpa variabel bantu).
	var a, b int = 15, 30

	p1, p2 := &a, &b

	*p1 = *p1 + *p2 // a sekarang jadi 45 (15 + 30)
	*p2 = *p1 - *p2 // b sekarang jadi 15 (45 - 30)
	*p1 = *p1 - *p2 // a sekarang jadi 30 (45 - 15)

	fmt.Printf("Sesudah: a = %d, b = %d\n", a, b)

}
