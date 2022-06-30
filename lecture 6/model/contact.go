package model

type Contact struct {
	contactID      string
	UniqueName     string
	Fname          string
	Lname          string
	IsActive       bool
	ContactDetails []ContactDetail
}

var AllContact []Contact

func CreateContacts(Fname, Lname, Type string, Number int) *Contact {
	var newcontact Contact

	newcontact.Fname = Fname
	newcontact.Lname = Lname
	newcontact.IsActive = true
	newcontact.ContactDetails = make([]ContactDetail, 0)
	CreateNewContactDetail(Type, Number)
	AllContact = append(AllContact, newcontact)
	return &newcontact
}
func (c *Contact) updateContact(ContactToBeupdated Contact, Fname, Lname, Type, ContactDetailID string, IsActive bool, Number int) {

	ContactToBeupdated.Fname = Fname
	ContactToBeupdated.Lname = Lname
	ContactToBeupdated.IsActive = IsActive
	ContactExist, LocationInArray := c.DoesContactDetailExist(ContactDetailID)
	if ContactExist == true {
		ContactToBeupdated.ContactDetails[LocationInArray].Type = Type
		ContactToBeupdated.ContactDetails[LocationInArray].Number = Number
	}

}

func (c *Contact) DoesContactDetailExist(ContactID string) (bool, int) {
	isContactExist := false
	isDataExist := -1
	for i := 0; i < len(c.ContactDetails); i++ {
		if ContactID == c.ContactDetails[i].ContactDetailID {
			isContactExist = true
			isDataExist = i
			break

		}

	}
	return isContactExist, isDataExist
}

func (u *User) DeleteContact(ContactToDelete string) {
	ContactExist, LocationInArray := u.DoesContactDetailExist(ContactToDelete)
	if ContactExist == true {
		u.Contacts[LocationInArray].IsActive = false
	}
}
