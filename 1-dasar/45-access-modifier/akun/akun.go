package akun

import "fmt"

type Akun struct {
	Norek   int
	Nasabah string
	saldo   int //private field diawali huruf kecil
}

func (a Akun) GetSaldo() {
	fmt.Println(a)
}

func (a *Akun) Setor(amount int) {
	a.saldo += amount
	fmt.Println("Berhasil menambah saldo")
}

func (a *Akun) Tarik(amount int) {
	if amount > a.saldo {
		fmt.Println("Saldo tidak mencukupi")
		return
	}
	a.saldo -= amount
	fmt.Println("Berhasil menarik saldo")
}
