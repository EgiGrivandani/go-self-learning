package main

import "fmt"

type Akun struct {
	Nama  string
	Saldo int
}

func (a Akun) Profile() {
	fmt.Printf("Akun: %s | Saldo: Rp %d\n", a.Nama, a.Saldo)
}

func (a *Akun) deposit(amount int) {
	a.Saldo += amount
	fmt.Printf("Deposit: %d | Saldo: Rp %d\n", amount, a.Saldo)
}

func (a *Akun) withdarw(amount int) bool {
	balance := a.Saldo
	if amount > balance {
		fmt.Printf("Saldo tidak cukup!")
		return false
	}

	a.Saldo -= amount
	fmt.Printf("Withdraw: %d | Saldo: Rp %d\n", amount, a.Saldo)
	return true
}

type TodoItem struct {
	Judul   string
	Selesai bool
}

func (t *TodoItem) Tandai(Newstatus bool) bool {
	if t.Selesai == Newstatus {
		fmt.Println("Tidak bisa mengubah status dengan status yang sama!")
		return false
	}
	t.Selesai = Newstatus
	fmt.Println("Berhasil mengubah status")
	return true
}

func (t *TodoItem) Info() {
	fmt.Printf("Judul: %s | Selesai %t\n", t.Judul, t.Selesai)
}

func main() {
	acc1 := Akun{"Egi", 100000}
	acc1.Profile()
	acc1.deposit(600000)
	acc1.withdarw(600000)
	acc1.withdarw(600000)

	//LATIHAN : Buat struct TodoItem dengan field Judul dan Selesai. Buat method Tandai() (pointer receiver) untuk mengubah Selesai menjadi true, dan method Info() untuk mencetak status.
	todo1 := TodoItem{"Springs", false}
	todo1.Tandai(true)
	todo1.Tandai(true)
	todo1.Info()
}
