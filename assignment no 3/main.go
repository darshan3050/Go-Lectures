package main

import (
	"model/model"
)

func main() {
	newbank1, _ := model.CreateNewBank("Axis Bank", "Axis")
	newbank2, _ := model.CreateNewBank("HDFC Bank", "HDFC")

	newUser1 := model.CreateNewCustomer("Darshan", "Naik")
	newUser2 := model.CreateNewCustomer("Rahul", "Nair")

	// fmt.Println(newbank2.BankId)
	// fmt.Println(newbank2.Abberivation)
	// fmt.Println(newbank2.Name)

	// fmt.Println(newUser1.Accounts)
	// fmt.Println(newUser1.CustomerId)
	// fmt.Println(newUser1.TotalBalance)
	// fmt.Println(newUser1.FirstName)
	// fmt.Println(newUser1.LastName)
	newUser1.CreateNewAccount(*newbank1)
	newUser2.CreateNewAccount(*newbank2)
	//fmt.Println(newUser1.Accounts[0].AccountNo)
	newUser2.CreateNewAccount(*newbank1)
	//fmt.Println(newUser2.Accounts[0].AccountNo)

	newUser2.DepositeCash(10000, 0)
	newUser2.DisplayTotalBalance(*newUser2)
	newUser2.DisplayIndivusalAccountBalance()
	newUser2.TransferAmmount(3000, 0, 1)
	newUser2.WidrawCash(3000, 0)
	newUser2.DisplayIndivusalAccountBalance()

}
