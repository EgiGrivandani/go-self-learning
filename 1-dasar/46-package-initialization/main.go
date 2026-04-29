package main

import "fmt"

func init() {
	fmt.Println("Pesan dari init 1")
}

func init() {
	fmt.Println("Pesan dari init 2")
}

func main() {
	fmt.Println("Pesan dari main")
}
