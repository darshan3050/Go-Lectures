package model

import (
	"fmt"

	"github.com/google/uuid"
)

type ContactDetail struct {
	ContactDetailID string
	Type            string
	Number          int
}

var DetailContact = make([]ContactDetail, 5)

func CreateNewContactDetail(Type string, Number int) (*ContactDetail, error) {

	var newContact ContactDetail
	var generateUUID = uuid.New()
	newContact.ContactDetailID = generateUUID.String()
	newContact.Type = Type
	newContact.Number = Number
	DetailContact = append(DetailContact, newContact)
	return &newContact, nil
}
func (u *User) ReadContactDetail() {

	fmt.Println(" User ContactDetails")
	for i := 0; i < len(u.Contacts); i++ {

		fmt.Println(u.Contacts[i].ContactDetails)

	}
	// fmt.Println(" User Id of InActive Users")
	// for i := 0; i < len(u.Contacts); i++ {
	// 	if !u.Contacts[i].isActive {
	// 		fmt.Println(u.UserID)
	// 	}
	// }
}
