package main

import "fmt"

func main() {
	var (
		umur     int  = 18
		isRegist bool = true
	)

	if umur >= 18 && isRegist {
		fmt.Println("Boleh ikut lomba")
	} else {
		fmt.Println("Salah 1 syarat tidak terpenuhi")
	}
}
