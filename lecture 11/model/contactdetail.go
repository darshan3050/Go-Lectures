package model

import (
	"fmt"

	"github.com/google/uuid"
)

type ContactDetail struct {
	ContactDetailID string
	Type            string
	Number          int
	IsActive        bool
}

var DetailContact = make([]ContactDetail, 5)

func CreateNewContactDetail(Type string, Number int) (*ContactDetail, error) {
	isContactExist := false
	for i := 0; i < len(DetailContact); i++ {
		if Number == DetailContact[i].Number || Type == DetailContact[i].Type {
			isContactExist = true
			break
		}
	}
	if isContactExist {
		return nil, fmt.Errorf("Contact Already Exist")
	}
	var newContact ContactDetail
	var generateUUID = uuid.New()
	newContact.ContactDetailID = generateUUID.String()
	newContact.Type = Type
	newContact.Number = Number
	newContact.IsActive = true
	DetailContact = append(DetailContact, newContact)
	return &newContact, nil
}
