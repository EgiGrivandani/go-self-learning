package account

import (
	"fmt"
	"math/rand"
)

type Account struct {
	Norek   int
	Nik     string
	Nasabah string
	Saldo   int
	Status  int
}

type ListAccount struct {
	List map[string]Account
}

func (l ListAccount) existNik(Nik string) bool {
	for _, v := range l.List {
		if v.Nik == Nik {
			return true
		}
	}
	return false
}

func (l ListAccount) existNorek(Norek int) bool {
	for _, v := range l.List {
		if v.Norek == Norek {
			return true
		}
	}
	return false
}

func (l ListAccount) GetId(nik string) {
	fmt.Println("====DATA Nasabah====")
	h := l.List[nik]
	fmt.Printf("%s - %d - %s -  %d - Rp %d\n", h.Nik, h.Norek, h.Nasabah, h.Status, h.Saldo)
}

func (l ListAccount) GetAll() {
	fmt.Println("====DATA ACCOUNT====")
	for _, h := range l.List {
		fmt.Printf("%s - %d - %s -  %d - Rp %d\n", h.Nik, h.Norek, h.Nasabah, h.Status, h.Saldo)
	}
}

func (l *ListAccount) Add(Nik, Nasabah string) bool {
	if l.existNik(Nik) {
		fmt.Println("Nik sudah terdaftar sebelumnya")
		return false
	}

	var norek int
	for {
		norek = rand.Intn(10000)

		if !l.existNorek(norek) {
			break
		}
	}

	l.List[Nik] = Account{
		Nik:     Nik,
		Norek:   norek,
		Nasabah: Nasabah,
		Saldo:   0,
		Status:  1,
	}
	return true
}

func (l *ListAccount) Blocked(nik string) bool {
	acc, ok := l.List[nik]
	if !ok {
		fmt.Println("Account tidak ditemukan")
		return false
	}

	if acc.Status == 0 {
		fmt.Println("Gagal, akun sudah di block sebelumnya")
		return false
	}

	acc.Status = 0
	l.List[nik] = acc

	fmt.Println("Berhasil memblokir Akun")
	return true
}
func (l *ListAccount) Active(nik string) bool {
	acc, ok := l.List[nik]
	if !ok {
		fmt.Println("Account tidak ditemukan")
		return false
	}

	if acc.Status == 1 {
		fmt.Println("Gagal, akun sudah aktif sebelumnya")
		return false
	}

	acc.Status = 1
	l.List[nik] = acc

	fmt.Println("Berhasil mengaktifkan Akun")
	return true
}
