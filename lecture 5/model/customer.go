package model

import (
	"fmt"

	"github.com/google/uuid"
)

type Customer struct {
	CustomerId   string
	FirstName    string
	LastName     string
	TotalBalance float64
	Accounts     []Account
}

func CreateNewCustomer(FirstName, LastName string) *Customer {
	var newCustomer Customer
	generateUUID := uuid.New()
	newCustomer.CustomerId = generateUUID.String()
	newCustomer.Accounts = make([]Account, 0)
	newCustomer.FirstName = FirstName
	newCustomer.LastName = LastName
	newCustomer.TotalBalance = 0
	return &newCustomer

}
func (c *Customer) CreateNewAccount(bank Bank) error {
	var isBankExist = bank.isBankExist(bank.Name)
	if !isBankExist {
		return fmt.Errorf("Bank Doesnot Exist")

	}
	var isAccountExist = c.isAccountExist(bank.Name)

	if isAccountExist == true {

		return fmt.Errorf("Account Already Exist")
	}
	// min := 3400000000
	// max := 3499999999
	// Account_Number :=

	newAccount := Create_Account(bank, 1000)
	c.Accounts = append(c.Accounts, *newAccount)
	c.UpdateTotalBalance()
	return fmt.Errorf("Account Created Sucessfully")
}

func (c *Customer) UpdateTotalBalance() {
	c.TotalBalance = 0
	for i := 0; i < len(c.Accounts); i++ {
		c.TotalBalance = c.TotalBalance + c.Accounts[i].AccountBalance
	}

}

func (c *Customer) isAccountExist(BankName string) bool {
	isAccountExist := false
	for i := 0; i < len(c.Accounts); i++ {
		if BankName == c.Accounts[i].AccountOfBank.Name {
			isAccountExist = true
			break

		}

	}
	return isAccountExist

}
func (c *Customer) DisplayTotalBalance(Customer) {
	c.UpdateTotalBalance()
	fmt.Println("Your Total Balance Ammount is : ", c.TotalBalance)
}

func (c *Customer) DisplayBalance() {

	for i := 0; i < len(c.Accounts); i++ {
		// 	fmt.Println("Balance of Account ", c.Accounts[i].AccountNo, "of Bank ", c.Accounts[i].AccountOfBank.Name, " is : Rs", c.Accounts[i].AccountBalance)
		c.Accounts[i].DisplayIndivusalAccountBalance()
	}
}

// func (c *Customer) DepositeCash(DepositeAmmount, Account_no int) {

// 	isAccountExist := false
// 	for i := 0; i < len(c.Accounts); i++ {
// 		if Account_no == c.Accounts[i].AccountNo {
// 			isAccountExist = true
// 			break

// 		}

// 	}
// 	if isAccountExist == true {
// 		fmt.Println("Previous Balance: Rs", c.Accounts[Account_no].AccountBalance)
// 		c.Accounts[Account_no].AccountBalance = c.Accounts[Account_no].AccountBalance + float64(DepositeAmmount)
// 		fmt.Println("New Account Balance : Rs", c.Accounts[Account_no].AccountBalance)
// 	}
// }

// func (c *Customer) TransferAmmount(TransferAmmount, Account_no1, Account_no2 int) {
// 	if c.Accounts[Account_no1].AccountBalance < float64(TransferAmmount) {
// 		fmt.Println("No Sufficient Funds")

// 	} else {
// 		c.Accounts[Account_no1].AccountBalance = c.Accounts[Account_no1].AccountBalance - float64(TransferAmmount)
// 		c.Accounts[Account_no2].AccountBalance = c.Accounts[Account_no2].AccountBalance + float64(TransferAmmount)
// 		fmt.Println("Transfer Sucessfull. Updated Balance is : Rs", c.Accounts[Account_no1].AccountBalance)
// 	}

// }
// func (c *Customer) WidrawCash(WidrawalAmmount, Account_no int) {
// 	if c.Accounts[Account_no].AccountBalance < float64(WidrawalAmmount) {
// 		fmt.Println("No Sufficient Funds")

// 	} else {
// 		fmt.Println("Previous Balance: Rs", c.Accounts[Account_no].AccountBalance)
// 		c.Accounts[Account_no].AccountBalance = c.Accounts[Account_no].AccountBalance - float64(WidrawalAmmount)
// 		fmt.Println("New Account Balance : Rs", c.Accounts[Account_no].AccountBalance)
// 	}
// }
func (c *Customer) findAccount(BankName string) (bool, int) {
	isAccountExist := false
	indexofAccount := -1
	for i := 0; i < len(c.Accounts); i++ {
		if BankName == c.Accounts[i].AccountOfBank.Name {
			isAccountExist = true
			break
		}
	}
	return isAccountExist, indexofAccount
}
func (c *Customer) Deposit(BankName string, Ammount float64) error {
	isAccountExist := false
	indexofAccount := -1
	isAccountExist, indexofAccount = c.findAccount(BankName)
	if !isAccountExist {

		return fmt.Errorf(("Bank Account Not Found"))

	}
	c.Accounts[indexofAccount].AccountBalance += Ammount
	c.UpdateTotalBalance()
	c.DisplayBalance()
	return nil

}

func (c *Customer) Widraw(BankName string, Ammount float64) error {
	isAccountExist := false
	indexofAccount := -1
	isAccountExist, indexofAccount = c.findAccount(BankName)
	if !isAccountExist {

		return fmt.Errorf(("Bank Account Not Found"))

	}
	if c.Accounts[indexofAccount].isSufficientBalance(Ammount) {

		return fmt.Errorf(("Insufficient Balance"))
	}

	c.Accounts[indexofAccount].AccountBalance -= Ammount
	c.UpdateTotalBalance()
	c.DisplayBalance()
	return nil
}
func (debitCustomer *Customer) TransferAmmount(creditCustomer *Customer, debitAccountBankName, creditAccountBankName string, Ammount float64) {
	err := debitCustomer.Widraw(debitAccountBankName, Ammount)
	if err != nil {
		fmt.Println(err)
		return
	}

	creditCustomer.Deposit(creditAccountBankName, Ammount)
	if err != nil {
		debitCustomer.Deposit(debitAccountBankName, Ammount)
		fmt.Println(err)
		return
	}

}
func (c *Customer) selfTransfer(debitAccountBankName, creditAccountBankName string, Ammount float64) {
	c.Widraw(debitAccountBankName, Ammount)
	c.Deposit(creditAccountBankName, Ammount)
}
