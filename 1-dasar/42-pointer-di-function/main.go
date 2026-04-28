package main

import "fmt"

func tambahSatu(number *int) {
	*number += 1
}

func tukarNilai(x, y *int) {
	*x, *y = *y, *x
}

type User struct {
	username string
	password string
}

func resetPassword(u *User) {
	u.password = "default123"
}

func main() {
	angka := 10
	fmt.Println("Sebelum:", angka)
	tambahSatu(&angka)
	fmt.Println("Sesudah:", angka)

	x, y := 100, 200
	fmt.Println("Sebelum: x=", x, "y=", y)
	tukarNilai(&x, &y)
	fmt.Println("Sesudah: x=", x, "y=", y)

	//LATIHAN : Buat function resetPassword(user *User) yang menerima pointer ke struct User dan mengubah field Password menjadi "default123". Buktikan bahwa data asli berubah.
	user := User{username: "egi", password: "123456"}
	fmt.Println("Sebelum Reset ", user)
	resetPassword(&user)
	fmt.Println("Sesudah Reset ", user)

}
