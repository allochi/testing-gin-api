package mocks

import (
	"cmp"
	"slices"
)

type Account struct {
	Address string
}

func GetAccounts() []Account {
	return []Account{
		{Address: "0x5678"},
		{Address: "0x1234"},
	}
}

func GetSortedAccounts() []Account {
	accounts := GetAccounts()
	slices.SortFunc(accounts, func(i, j Account) int {
		return cmp.Compare(i.Address, j.Address)
	})
	return accounts
}
