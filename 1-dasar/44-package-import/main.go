package main

import (
	"44-package-import/account"
)

func main() {
	list := account.ListAccount{
		List: make(map[string]account.Account),
	}

	list.Add("3214122809980001", "Egi grivandani")
	list.Add("3214122809980003", "Budi dorks")

	list.Blocked("3214122809980001")
	list.GetId("3214122809980001")
}
