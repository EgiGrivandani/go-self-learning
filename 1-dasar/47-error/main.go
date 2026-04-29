package main

import (
	"errors"
	"fmt"
)

func Pembagian(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Tidak bisa di bagi dengan 0")
	}
	return a / b, nil
}

func tarikUang(saldo, jumlah int) (int, error) {
	if jumlah > saldo {
		return saldo, errors.New("jumlah lebih besar dari saldo.")
	}

	saldo -= jumlah
	return saldo, nil

}

func main() {

	hasil, err := Pembagian(10, 0)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(hasil)
	}

	//LATIHAN : Buat function tarikUang(saldo, jumlah int) (int, error) yang mengembalikan sisa saldo dan error jika jumlah lebih besar dari saldo.
	res, err := tarikUang(100, 170)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(res)
	}

}
