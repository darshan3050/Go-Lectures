package main

import (
	"fmt"
	"model/model"
)

func main() {
	admin := model.CreateAdmin("Darshan", "Naik")
	user1 := model.CreateNewUser("Prathamesh", "Naik", false, *admin)
	user2 := model.CreateNewUser("Rahul", "Nair", false, *admin)

	fmt.Println("\n\n Admin Commands")
	admin.ReadUsers()

	admin.UpdateUser(true, true, user1)

	admin.DeleteUser(user2)
	admin.ReadUsers()

	fmt.Println("\n\n\n\n\n Contact Commands")

	user1.CreatenewContact("Rahul", "nair", "home", 9876549875)
	user1.ReadContacts()
	user1.CreatenewContact("Prathamesh", "naik", "home", 9876549875)

	user1.ReadContacts()
	user1.DeleteContact("prathamesh")

	user1.ReadContactDetail()
}
