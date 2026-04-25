package main

import "fmt"

func main() {
	//LATIHAN : Buat program yang mencetak angka 1-20, tapi skip angka kelipatan 3 (pakai continue) dan berhenti jika angka sudah lebih dari 15 (pakai break).
	breakContinue(1)
}

func breakContinue(n int) {
	for {
		if n > 15 {
			fmt.Println("Stop")
			break
		}

		if n%3 == 0 {
			fmt.Println("Skip")
			n++
			continue
		}
		fmt.Println("Angka ke-", n)
		n++
	}
}
