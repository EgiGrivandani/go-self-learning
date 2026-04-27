package main

import "fmt"

type Hewan interface {
	suara() string
	jenis() string
}

type Kucing struct{}

func (k Kucing) suara() string {
	return "Meong"
}

func (k Kucing) jenis() string {
	return "Mamalia"
}

type Anjing struct{}

func (k Anjing) suara() string {
	return "Guk Guk"
}

func (k Anjing) jenis() string {
	return "Mamalia"
}

type Burung struct{}

func (k Burung) suara() string {
	return "Cuit Cuit"
}

func (k Burung) jenis() string {
	return "Aves"
}

func main() {
	//LATIHAN : Buat interface Hewan dengan method Suara() string dan Jenis() string. Implementasikan untuk struct Kucing, Anjing, dan Burung. Buat slice []Hewan dan cetak suara semua hewan.
	daftarHewan := []Hewan{Kucing{}, Anjing{}, Burung{}}

	for _, h := range daftarHewan {
		fmt.Printf("Suara %s | Jenis %s \n", h.suara(), h.jenis())
	}
}
