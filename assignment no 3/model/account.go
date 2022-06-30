package model

import (
	"math/rand"
)

type Account struct {
	AccountNo      int
	AccountOfBank  Bank
	AccountBalance float64
}

func Create_Account(AccountOfBank Bank, AccountBalance float64) *Account {
	var newAccount Account
	newAccount.AccountBalance = 1000
	min := 3400000000
	max := 3499999999

	generateUUID := rand.Intn(max-min) + min
	newAccount.AccountNo = generateUUID
	newAccount.AccountOfBank = AccountOfBank
	return &newAccount
}
