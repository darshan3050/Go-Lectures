package model

import "fmt"

type Contact struct {
	contactID      string
	Fname          string
	Lname          string
	isActive       bool
	ContactDetails []ContactDetail
}

var AllContacts []Contact

func (u *User) CreateContacts(Fname, Lname, Type string, Number int) *Contact {
	var newcontact Contact

	newcontact.Fname = Fname
	newcontact.Lname = Lname
	newcontact.isActive = true
	Contact, _ := CreateNewContactDetail(Type, Number)
	newcontact.ContactDetails = append(newcontact.ContactDetails, *Contact)
	AllContacts = append(AllContacts, newcontact)
	// u.Contacts = append(u.Contacts, newcontact)
	return &newcontact
}
func (u *User) UpdateContact(ContactToBeupdated []Contact, contactname, Fname, Lname, Type string, IsActive bool, Number int) {
	isContact_Exist, Index := u.DoesContactDetailExist(contactname)
	if isContact_Exist {
		u.Contacts[Index].Fname = Fname
		u.Contacts[Index].Lname = Lname
		u.Contacts[Index].isActive = IsActive
		detailExist, index := u.DoesContactDetailExist(Type)
		if detailExist {
			u.Contacts[Index].ContactDetails[index].Type = Type
			u.Contacts[Index].ContactDetails[index].Number = Number
		}

	}

}

func (c *Contact) DoesContactDetailExist(contactName string) (bool, int) {
	isContactExist := false
	isDataExist := -1
	for i := 0; i < len(AllContacts); i++ {
		if contactName == AllContacts[i].Fname {
			isContactExist = true
			isDataExist = i
			break

		}

	}
	return isContactExist, isDataExist
}

func (u *User) DeleteContact(ContactToDelete string) {
	ContactExist, LocationInArray := u.DoesContactDetailExist(ContactToDelete)
	if ContactExist {
		u.Contacts[LocationInArray].isActive = false
	}
}
func (u *User) ReadContacts() {

	fmt.Println(" Contacts List")
	for i := 0; i < len(AllContacts); i++ {
		if u.Contacts[i].isActive {
			fmt.Println(u.Contacts[i])
		}
	}

}
