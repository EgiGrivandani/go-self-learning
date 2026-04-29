package account

import (
	"fmt"
)

type Account struct {
	Norek   int
	Nik     string
	Nasabah string
	Saldo   int
	Status  int
}

type ListAccount struct {
	List map[int]Account
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

func (l ListAccount) FindNorek(norek int) (Account, bool) {
	acc, ok := l.List[norek]
	return acc, ok
}

func (l ListAccount) GetId(norek int) {
	fmt.Println("====DATA Nasabah====")
	h := l.List[norek]
	fmt.Printf("%d - %s - %s -  %d - Rp %d\n", h.Norek, h.Nik, h.Nasabah, h.Status, h.Saldo)
}

func (l ListAccount) GetAll() {
	fmt.Println("====DATA ACCOUNT====")
	for _, h := range l.List {
		fmt.Printf("%d - %s - %s -  %d - Rp %d\n", h.Norek, h.Nik, h.Nasabah, h.Status, h.Saldo)
	}
}

func (l *ListAccount) Add(Nik string, Norek int, Nasabah string) bool {
	if l.existNik(Nik) {
		fmt.Println("Nik sudah terdaftar sebelumnya")
		return false
	}

	if l.existNorek(Norek) {
		fmt.Println("No Rekening sudah terdaftar sebelumnya")
		return false
	}

	l.List[Norek] = Account{
		Nik:     Nik,
		Norek:   Norek,
		Nasabah: Nasabah,
		Saldo:   0,
		Status:  1,
	}
	return true
}

func (l *ListAccount) Blocked(norek int) bool {
	acc, ok := l.List[norek]
	if !ok {
		fmt.Println("Account tidak ditemukan")
		return false
	}

	if acc.Status == 0 {
		fmt.Println("Gagal, akun sudah di block sebelumnya")
		return false
	}

	acc.Status = 0
	l.List[norek] = acc

	fmt.Println("Berhasil memblokir Akun")
	return true
}
func (l *ListAccount) Active(norek int) bool {
	acc, ok := l.List[norek]
	if !ok {
		fmt.Println("Account tidak ditemukan")
		return false
	}

	if acc.Status == 1 {
		fmt.Println("Gagal, akun sudah aktif sebelumnya")
		return false
	}

	acc.Status = 1
	l.List[norek] = acc

	fmt.Println("Berhasil mengaktifkan Akun")
	return true
}
