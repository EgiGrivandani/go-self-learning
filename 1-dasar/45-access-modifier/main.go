package main

import (
	"45-access-modifier/akun"
	"fmt"
)

func main() {
	akun := akun.Akun{}
	akun.Norek = 456743
	akun.Nasabah = "Egi"

	akun.Setor(1000000)
	akun.Tarik(20000)
	fmt.Println(akun)
}
