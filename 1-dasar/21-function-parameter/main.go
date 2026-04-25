package main

import "fmt"

func main() {
	buahBuahan := []string{"Semangka", "Mangga", "Pepaya"}
	cetakDaftar("Daftar Buah", buahBuahan)

}

func cetakDaftar(judul string, items []string) {
	fmt.Println("====", judul, "====")
	for _, item := range items {
		fmt.Println("- ", item)
	}

}
