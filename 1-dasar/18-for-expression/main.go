package main

import "fmt"

func main() {
	//LATIHAN : Buat program yang mencetak angka 1 sampai 20, tapi hanya cetak angka yang genap saja (gunakan if di dalam for).

	fmt.Println("For while")
	forWhile(1)
	fmt.Println("\nFor infinite")
	forInfinite(1)
	fmt.Println("\nFor classic")
	forClassic(1)
}

func forWhile(n int) {
	for n <= 20 {
		if n%2 == 0 {
			fmt.Println("Even number", n)
		}
		n++
	}
}

func forInfinite(n int) {
	for {
		if n%2 == 0 {
			fmt.Println("Even number", n)
		}
		if n > 20 {
			break
		}
		n++
	}
}

func forClassic(n int) {
	for i := n; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println("Even number", i)
		}
	}
}
