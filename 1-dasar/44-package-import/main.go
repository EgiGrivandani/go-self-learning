package main

import (
	"44-package-import/account"
	"44-package-import/transaction"
)

func main() {
	list := account.ListAccount{
		List: make(map[int]account.Account),
	}

	list.Add("3214122809980001", 8709210, "Egi grivandani")
	list.Add("3214122809980003", 8701231, "Budi dorks")
	list.GetId(8709210)

	Trx := transaction.List{
		List: make(map[int]transaction.Transaksi),
	}

	Trx.Deposit(&list, 8709210, 190000)
	list.GetId(8709210)
	Trx.Withdraw(&list, 8709210, 100000)
	list.GetId(8709210)
	Trx.Transfer(&list, 8709210, 8701231, 110000)
	list.GetId(8709210)
	list.GetId(8701231)
}
