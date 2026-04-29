package transaction

import (
	"44-package-import/account"
	"fmt"
)

type Transaksi struct {
	Norek  int
	Amount int
	Type   string
}

type List struct {
	List map[int]Transaksi
}

func onlyActive(acc *account.ListAccount, norek int) (account.Account, error) {
	a, ok := acc.FindNorek(norek)
	if !ok {
		return a, fmt.Errorf("account tidak ditemukan")
	}
	if a.Status == 0 {
		return a, fmt.Errorf("akun diblokir")
	}
	return a, nil
}

func (h *List) Deposit(acc *account.ListAccount, norek, amount int) {
	fmt.Println("====DEPOSIT====")
	a, err := onlyActive(acc, norek)
	if err != nil {
		fmt.Println(err)
		return
	}

	a.Saldo += amount
	acc.List[norek] = a

	h.List[norek] = Transaksi{
		Norek:  norek,
		Amount: amount,
		Type:   "Deposit",
	}
	fmt.Println("Berhasil deposit sebesar Rp.", amount)
}

func (h *List) Withdraw(acc *account.ListAccount, norek, amount int) {
	fmt.Println("====WITDRAW====")
	a, err := onlyActive(acc, norek)
	if err != nil {
		fmt.Println(err)
		return
	}

	if amount > a.Saldo {
		fmt.Println("Saldo tidak mencukupi untuk penarikan ini")
		return
	}

	a.Saldo -= amount
	acc.List[norek] = a

	h.List[norek] = Transaksi{
		Norek:  norek,
		Amount: amount,
		Type:   "Withdraw",
	}
	fmt.Println("Berhasil withdraw sebesar Rp.", amount)
}

func (h *List) Transfer(acc *account.ListAccount, from, to, amount int) {
	x, err := onlyActive(acc, from)
	if err != nil {
		fmt.Println("Pengirim : ", err)
		return
	}

	y, err := onlyActive(acc, to)
	if err != nil {
		fmt.Println("Penerima : ", err)
		return
	}

	if amount > x.Saldo {
		fmt.Println("Saldo tidak mencukupi dalam transfer ini")
		return
	}

	x.Saldo -= amount
	acc.List[from] = x

	y.Saldo += amount
	acc.List[to] = y

	h.List[from] = Transaksi{
		Norek:  from,
		Amount: amount,
		Type:   "Transfer",
	}

	fmt.Println("Transfer berhasil!!")
}
