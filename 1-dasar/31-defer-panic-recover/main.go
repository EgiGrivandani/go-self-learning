package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApp() {
	defer logging()
	fmt.Println("Run App")
}

func samplePanic() {
	fmt.Println("\n=== Panic ===")
	fmt.Println("Sebelum panic")
	panic("Terjadi error yang tidak bisa ditangani!")
}

func sampleRecover() {
	fmt.Println("\n=== Recover ===")

	defer func() {
		message := recover()
		if message != nil {
			fmt.Println("Panic ditangkap:", message)
			fmt.Println("Program tetap berjalan!")
		}
	}()
	fmt.Println("Sebelum panic")
	panic("Error: sesuatu yang buruk terjadi")

	//fmt.Println("Pesan ini tidak akan di munculkan, karena sudah panic")
}

func main() {
	runApp()

	//samplePanic()

	sampleRecover()
	fmt.Println("Pesan ini tetap akan di munculkan, walaupun sudah panic")
}
