package main

import (
	"fmt"
	"model/model"
)

func main() {
	newuser1 := model.CreateUser("Darshan", "Naik", true)
	newuser2 := model.CreateUser("Prathamesh", "Naik", false)
	newuser3 := model.CreateUser("Rahul", "Nair", false)

	fmt.Println(newuser1)
	fmt.Println(newuser2)
	fmt.Println(newuser3)
	newuser2.CreatenewContact("FnameTest", "lnameTest", "Home", 9876543210)
	newuser2.ReadUsers()

	newuser1.CreatenewContact("FnameTest", "lnameTest", "Home", 9876543210)
	newuser1.ReadUsers()
}
