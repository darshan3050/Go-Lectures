package model

import (
	"fmt"

	"github.com/google/uuid"
)

type Bank struct {
	BankId       string
	Name         string
	Abberivation string
}

var AllBanks = make([]Bank, 5, 5)

func CreateNewBank(Name, Abberivation string) (*Bank, error) {
	isBankExist := false
	for i := 0; i < len(AllBanks); i++ {
		if Name == AllBanks[i].Name || Abberivation == AllBanks[i].Abberivation {
			isBankExist = true
			break
		}
	}
	if isBankExist {
		return nil, fmt.Errorf("bank Already Exist")
	}

	var newbank Bank
	var generateUUID = uuid.New()
	newbank.Name = Name
	newbank.BankId = generateUUID.String()
	newbank.Abberivation = Abberivation
	AllBanks = append(AllBanks, newbank)
	return &newbank, nil

}
func (b *Bank) isBankExist(BankName string) bool {
	isBankExist := false
	for i := 0; i < len(AllBanks); i++ {
		if BankName == AllBanks[i].Name {
			isBankExist = true
			break
		}

	}
	return isBankExist
}
