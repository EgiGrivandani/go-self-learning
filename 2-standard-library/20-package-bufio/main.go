package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=====Baca baris per baris======")
	text := "Halo dunia\nBelajar Go itu seru"
	scan := bufio.NewScanner(strings.NewReader(text))

	total := 0
	for scan.Scan() {
		total += 1
		fmt.Println(scan.Text())
	}
	fmt.Println(total)

	//LATIHAN : Buat program yang membaca teks multi-baris, hitung jumlah baris, kata, dan karakter (seperti perintah wc di Linux).
	input := `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
			Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
			Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.

			Nisi ut aliquip ex ea commodo consequat.
			Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore.
			Eu fugiat nulla pariatur.`

	var totalBaris, totalKata, totalKarakter int

	// Hitung baris
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		totalBaris++
	}

	// Hitung kata
	scanner2 := bufio.NewScanner(strings.NewReader(input))
	scanner2.Split(bufio.ScanWords)
	for scanner2.Scan() {
		totalKata++
	}

	// Hitung karakter (rune)
	scanner3 := bufio.NewScanner(strings.NewReader(input))
	scanner3.Split(bufio.ScanRunes)
	for scanner3.Scan() {
		ch := scanner3.Text()
		if ch != " " && ch != "\n" && ch != "\t" {
			totalKarakter++
		}
	}

	totalByte := len(input)

	fmt.Println("Baris     :", totalBaris)
	fmt.Println("Kata      :", totalKata)
	fmt.Println("Karakter  :", totalKarakter)
	fmt.Println("Byte      :", totalByte)

}
