package model

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	UserID   string `json:"uid"`
	UniqueID string `json:"username"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isadmin"`
	IsActive bool   `json:"isactive"`
	Contacts []Contact
}

var AllUser []User

func CreateUser(UniqueID, Fname, Lname, Password string, Isadmin bool) *User {
	var newUser User
	generateUUID := uuid.New()
	newUser.UserID = generateUUID.String()
	newUser.UniqueID = UniqueID
	newUser.Fname = Fname
	newUser.Lname = Lname
	newUser.Password = Password
	newUser.IsAdmin = Isadmin
	newUser.IsActive = true
	newUser.Contacts = make([]Contact, 0)
	AllUser = append(AllUser, newUser)
	return &newUser
}
func (u *User) CreatenewContact(UniqueName, Fname, Lname, Type string, Number int) {

	newContact := CreateContacts(UniqueName, Fname, Lname, Type, Number)

	u.Contacts = append(u.Contacts, *newContact)
}

func (u *User) CreateNewUser(IsAdmin User, UniqueID, Fname, Lname, Password string, Isadmin bool) error {
	if !IsAdmin.IsAdmin {

		return fmt.Errorf("sou need Admin Previlages To Perform this task")
	} else {
		CreateUser(UniqueID, Fname, Lname, Password, Isadmin)
	}
	return fmt.Errorf("User Added to database")
}
func (u *User) ReadUsers(IsAdmin User) {

	if !IsAdmin.IsAdmin {

		fmt.Println("You need Admin Previlages To Perform this task")
	} else {

		fmt.Println(" User Id of Active Users")
		for i := 0; i < len(u.Contacts); i++ {
			if u.Contacts[i].IsActive {
				fmt.Println(u.UserID)
			}
		}

	}
}
func (u *User) UpdateUser(AdminUpdate, ActiveUpdate bool, IsAdmin, User User) {
	if !IsAdmin.IsAdmin {

		fmt.Println("You need Admin Previlages To Perform this task")
	} else {

		User.IsAdmin = AdminUpdate
		User.IsActive = ActiveUpdate
	}
}
func (u *User) DeleteUser(User *User, IsAdmin User) {
	if !IsAdmin.IsAdmin {

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
