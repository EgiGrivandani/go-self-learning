package main

import "fmt"

func main() {
	l, k := HitungLuasPersegiPanjang(10, 5)
	fmt.Println("Luas:", l)
	fmt.Println("Keliling:", k)

	fmt.Println(cekUmur(-1))
	fmt.Println(cekUmur(15))

	//Latihan : hitungBMI(berat, tinggiCm float64) dengan named return (bmi float64, kategori string). Hitung BMI (berat / tinggiM²) dan tentukan kategorinya: "Kurus" (<18.5), "Normal" (18.5-24.9), "Gemuk" (>=25).
	berat := 70.0
	tinggi := 175.0
	bmi, kategori := hitungBMI(berat, tinggi)
	fmt.Println("BMI:", bmi)
	fmt.Println("Kategori:", kategori)
}

func HitungLuasPersegiPanjang(panjang, lebar float64) (luas, keliling float64) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return

}

func cekUmur(umur int) (pesan string) {
	if umur < 0 {
		return "Umur tidak valid" // return eksplisit juga boleh
	}
	if umur >= 17 {
		pesan = "Boleh buat SIM"
	} else {
		pesan = "Belum boleh buat SIM"
	}
	return
}

func hitungBMI(berat, tinggiCm float64) (bmi float64, kategori string) {
	tinggiM := tinggiCm / 100
	bmi = berat / (tinggiM * tinggiM)
	if bmi < 18.5 {
		kategori = "Kurus"
	} else if bmi < 24.9 {
		kategori = "Normal"
	} else {
		kategori = "Gemuk"
	}
	return
}
