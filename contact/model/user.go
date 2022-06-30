package model

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	UserID   string
	Fname    string
	Lname    string
	IsAdmin  bool
	IsActive bool
	Contacts []Contact
}

func CreateAdmin(Fname, Lname string) *User {
	var newUser User
	generateUUID := uuid.New()
	newUser.UserID = generateUUID.String()
	newUser.Fname = Fname
	newUser.Lname = Lname
	newUser.IsAdmin = true
	newUser.IsActive = true
	newUser.Contacts = make([]Contact, 0)
	AllUser = append(AllUser, newUser)
	return &newUser
}

var AllUser []User

func CreateUser(Fname, Lname string, Isadmin bool) *User {

	var newUser User
	generateUUID := uuid.New()
	newUser.UserID = generateUUID.String()
	newUser.Fname = Fname
	newUser.Lname = Lname
	newUser.IsAdmin = Isadmin
	newUser.IsActive = true
	newUser.Contacts = make([]Contact, 0)

	return &newUser
}
func (u *User) CreatenewContact(Fname, Lname, Type string, Number int) {

	newContact := u.CreateContacts(Fname, Lname, Type, Number)

	u.Contacts = append(u.Contacts, *newContact)
}

func CreateNewUser(Fname, Lname string, Role bool, AdminCheck User) *User {
	if !AdminCheck.IsAdmin {

		fmt.Println("you need admin previlages to perform this task")
		return nil
	} else {
		newUser := CreateUser(Fname, Lname, Role)
		fmt.Println("User Added to database")
		AllUser = append(AllUser, *newUser)
		return newUser
	}

}
func (u *User) ReadUsers() {

	if !u.IsAdmin {

		fmt.Println("You need Admin Previlages To Perform this task    // Read User", u.IsAdmin)
	} else {

		fmt.Println(" User Id of Active Users")
		for i := 0; i < len(AllUser); i++ {
			if AllUser[i].IsActive {
				fmt.Println(AllUser[i])
			}
		}
		// fmt.Println(" User Id of InActive Users")
		// for i := 0; i < len(u.Contacts); i++ {
		// 	if !u.Contacts[i].isActive {
		// 		fmt.Println(u.UserID)
		// 	}
		// }
	}
}
func (u *User) UpdateUser(AdminUpdate, ActiveUpdate bool, User *User) {

	if !u.IsAdmin {

		fmt.Println("You need Admin Previlages To Perform this task")
	} else {
		isUserExist, indexOfUser := DoesUserExist(User)
		if isUserExist {
			AllUser[indexOfUser].IsActive = ActiveUpdate
			AllUser[indexOfUser].IsAdmin = AdminUpdate

		}
	}
}
func (u *User) DeleteUser(User *User) {

	if !u.IsAdmin {

		fmt.Println("You need Admin Previlages To Perform this task delete user")
	} else {
		isUserExist, indexOfUser := DoesUserExist(User)
		if isUserExist {
			AllUser[indexOfUser].IsActive = false

		}
	}

}
func DoesUserExist(User *User) (bool, int) {
	isUserExist := false
	isDataExist := -1
	for i := 0; i < len(AllUser); i++ {
		if AllUser[i].UserID == User.UserID {
			isUserExist = true
			isDataExist = i
			break

		}

	}
	return isUserExist, isDataExist
}
func (u *User) DoesContactDetailExist(ContactID string) (bool, int) {

	isContactExist := false
	isDataExist := -1
	for i := 0; i < len(u.Contacts); i++ {
		if ContactID == u.Contacts[i].contactID {
			isContactExist = true
			isDataExist = i
			break

		}

	}
	return isContactExist, isDataExist
}
