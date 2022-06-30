package card

type Card struct {
	CardValue int
}

func CreateNewCard(value int) *Card {
	var c Card
	c.CardValue = value
	return &c
}
