package main

import "fmt"

type mahaiswa struct {
	Nama    string
	NIM     string
	Jurusan string
	Ipk     float64
}

type produk struct {
	Nama  string
	Harga int
	stok  int
}

func main() {
	data := mahaiswa{
		Nama:    "egi",
		NIM:     "2024001",
		Jurusan: "Informatika",
		Ipk:     3.6,
	}

	fmt.Println(data)

	fmt.Println()

	mhs2 := mahaiswa{"Andi Pratama", "2024002", "Teknik Sipil", 3.50}
	fmt.Println(mhs2.Nama)
	fmt.Println(mhs2.NIM)
	fmt.Println(mhs2.Jurusan)
	fmt.Println(mhs2.Ipk)

	fmt.Println()

	//LATIHAN : Buat struct Produk dengan field: Nama, Harga, dan Stok. Buat 3 produk, simpan dalam slice, lalu loop dan cetak total harga semua produk (harga × stok).
	produk1 := produk{"sabun", 5000, 100}
	produk2 := produk{"pasta gigi", 35000, 50}
	produk3 := produk{"sampo", 12000, 10}

	daftarProduk := []produk{produk1, produk2, produk3}

	totalHarga := 0
	daftarTotal := make(map[string]int)
	for _, v := range daftarProduk {
		subtotal := v.Harga * v.stok
		totalHarga += subtotal

		daftarTotal[v.Nama] = subtotal
	}
	fmt.Println("Daftar pesanan ", daftarTotal)
	fmt.Println("Total harga ", totalHarga)
}
