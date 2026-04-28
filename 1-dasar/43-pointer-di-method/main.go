package main

import "fmt"

type Item struct {
	Nama  string
	Harga int
}

type Keranjang struct {
	Items []Item
}

func (k *Keranjang) Tambah(item Item) {
	k.Items = append(k.Items, item)

}

func (k *Keranjang) Hapus(nama string) {
	for i, item := range k.Items {
		if item.Nama == nama {
			k.Items = append(k.Items[:i], k.Items[i+1:]...)
			return
		}
	}
}

func (k Keranjang) Total() int {
	total := 0
	for _, item := range k.Items {
		total += item.Harga
	}
	return total
}

func (k Keranjang) Info() {
	fmt.Println("====Keranjang Belanja====")
	for i, item := range k.Items {
		fmt.Printf("  %d. %s - Rp %d\n", i+1, item.Nama, item.Harga)
	}
	fmt.Printf("  Total: Rp %d\n", k.Total())
}

type Account struct {
	Norek   int
	Nasabah string
	Saldo   int
}

type Log struct {
	Logs []Transaction
}

type Transaction struct {
	Norek  int
	Type   string
	Amount int
}

func (a Account) InfoSaldo() {
	fmt.Println("====Info saldo====")
	fmt.Printf("  %d - %s - Rp %d\n", a.Norek, a.Nasabah, a.Saldo)
}

func (a *Account) Deposit(amount int, l *Log) {
	a.Saldo += amount
	l.InsertLogs(a.Norek, "Deposit", amount)
	fmt.Printf(" berhasil menambah saldo sebesar - Rp %d\n", amount)
}

func (a *Account) Withdraw(amount int, l *Log) bool {
	if amount > a.Saldo {
		fmt.Println("Saldo tidak mencukupi")
		return false
	}
	a.Saldo -= amount
	l.InsertLogs(a.Norek, "Withdraw", amount)
	fmt.Printf(" berhasil withdraw sebesar - Rp %d\n", amount)
	return true
}

func (l *Log) InsertLogs(norek int, t string, amount int) {
	l.Logs = append(l.Logs, Transaction{
		Norek:  norek,
		Type:   t,
		Amount: amount,
	})
}

func (l Log) LogTransaksi() {
	fmt.Println("====Log transaksi====")
	for i, h := range l.Logs {
		fmt.Printf(" %d. %d - %s - Rp %d\n", i+1, h.Norek, h.Type, h.Amount)
	}
}

func main() {
	keranjang := Keranjang{}

	keranjang.Tambah(Item{Nama: "Buku Go", Harga: 85000})
	keranjang.Tambah(Item{Nama: "Keyboard", Harga: 350000})
	keranjang.Tambah(Item{Nama: "Mouse", Harga: 150000})
	keranjang.Info()

	fmt.Println("\nHapus Mouse...")
	keranjang.Hapus("Mouse")
	keranjang.Info()

	// fmt.Println("\n")
	//LATIHAN : Buat struct Tabungan dengan method pointer receiver Setor() dan Tarik(). Pastikan Tarik() tidak memperbolehkan saldo menjadi negatif.
	logData := Log{}
	nasabah := Account{Norek: 41131231, Nasabah: "Egi Grivandani", Saldo: 0}

	nasabah.Deposit(500000, &logData)
	nasabah.Deposit(1500000, &logData)
	nasabah.Withdraw(500000, &logData)
	nasabah.InfoSaldo()
	logData.LogTransaksi()
}
