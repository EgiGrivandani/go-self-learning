package main

import (
	"fmt"
	"slices"
	"sort"
)

type Produk struct {
	Nama   string
	Harga  int
	Rating float64
}

func main() {
	products := []Produk{
		{
			Nama:   "Produk A",
			Harga:  10000,
			Rating: 4.5,
		},
		{
			Nama:   "Produk B",
			Harga:  20000,
			Rating: 4.8,
		},
		{
			Nama:   "Produk C",
			Harga:  15000,
			Rating: 4.2,
		},
	}

	fmt.Println("======GO VEWRSI LAMA============")
	fmt.Println("\n asc harga")
	sort.Slice(products, func(i, j int) bool {
		return products[i].Harga < products[j].Harga
	})
	for _, s := range products {
		fmt.Printf("  %s: %d\n", s.Nama, s.Harga)
	}

	fmt.Println("\n desc rating")
	sort.Slice(products, func(i, j int) bool {
		return products[i].Rating > products[j].Rating
	})
	for _, s := range products {
		fmt.Printf("  %s: %.1f\n", s.Nama, s.Rating)
	}

	fmt.Println("\n======GO VEWRSI BARU============")
	slices.SortFunc(products, func(a, b Produk) int {
		if a.Rating > b.Rating {
			return -1
		}

		if a.Rating < b.Rating {
			return 1
		}

		if a.Harga < b.Harga {
			return -1
		}
		if a.Harga > b.Harga {
			return 1
		}
		return 0
	})
	for _, s := range products {
		fmt.Printf("  %s: %d\n", s.Nama, s.Harga)
	}
}
