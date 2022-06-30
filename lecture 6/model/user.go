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

	newContact := CreateContacts(Fname, Lname, Type, Number)

	u.Contacts = append(u.Contacts, *newContact)
}

func (u *User) CreateNewUser(Fname, Lname string, Role, Isadmin bool) error {
	if !Role == true {

		return fmt.Errorf("You need Admin Previlages To Perform this task")
	} else {
		CreateUser(Fname, Lname, Isadmin)
	}
	return fmt.Errorf("User Added to database")
}
func (u *User) ReadUsers() {

	if u.IsAdmin == false {

		fmt.Println("You need Admin Previlages To Perform this task")
	} else {

		fmt.Println(" User Id of Active Users")
		for i := 0; i < len(u.Contacts); i++ {
			if u.Contacts[i].isActive == true {
				fmt.Println(u.UserID)
			}
		}
		fmt.Println(" User Id of InActive Users")
		for i := 0; i < len(u.Contacts); i++ {
			if u.Contacts[i].isActive == false {
				fmt.Println(u.UserID)
			}
		}
	}
}
func (u *User) UpdateUser(AdminUpdate, ActiveUpdate bool, User *User) {
	if u.IsAdmin == false {

		fmt.Println("You need Admin Previlages To Perform this task")
	} else {

		User.IsAdmin = AdminUpdate
		User.IsActive = ActiveUpdate
	}
}
func (u *User) DeleteUser(User *User) {
	if u.IsAdmin == false {

		fmt.Println("You need Admin Previlages To Perform this task")
	} else {
		User.IsActive = false
	}

}
func (c *User) DoesContactDetailExist(ContactID string) (bool, int) {
	isContactExist := false
	isDataExist := -1
	for i := 0; i < len(c.Contacts); i++ {
		if ContactID == c.Contacts[i].contactID {
			isContactExist = true
			isDataExist = i
			break

		}

	}
	return isContactExist, isDataExist
}
