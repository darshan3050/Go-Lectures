package deck

import (
	"fmt"
	"math/rand"
	"module/card"
	"time"
)

type Deck struct {
	cardpack1 [13]card.Card
	cardpack2 [13]card.Card
	cardpack3 [13]card.Card
	cardpack4 [13]card.Card
}

func CreateNewDeck() *Deck {
	var d Deck
	for i := 0; i < 13; i++ {
		d.cardpack1[i] = *card.CreateNewCard(i + 1)
		d.cardpack2[i] = *card.CreateNewCard(i + 1)
		d.cardpack3[i] = *card.CreateNewCard(i + 1)
		d.cardpack4[i] = *card.CreateNewCard(i + 1)
	}

	return &d
}
func (d *Deck) DrawCard() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 4
	pack := (rand.Intn(max-min+1) + min)

	switch pack {
	case 1:
		c := DrawRandomCard(d.cardpack1)
		fmt.Println("\nTree Set")
		return c
	case 2:
		c := DrawRandomCard(d.cardpack2)
		fmt.Println("\nSpade Set")
		return c
	case 3:
		c := DrawRandomCard(d.cardpack3)
		fmt.Println("\nHeart Set")
		return c
	case 4:
		c := DrawRandomCard(d.cardpack4)
		fmt.Println("\nSquare Set ")
		return c
	default:
		return 0
	}
	// rand.Seed(time.Now().UnixNano())
	// min := 1
	// max := 4
	// fmt.Println(rand.Intn(max-min+1) + min)
	// randomIndex := rand.Intn(len(in))
	//rand.Intn(max - min) + min
}

func DrawRandomCard(pack [13]card.Card) int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 12
	randomIndex := (rand.Intn(max-min+1) + min)

	if pack[randomIndex].CardValue == 0 {
		DrawRandomCard(pack)
		return 0000
	} else {
		AddCardValue := pack[randomIndex].CardValue
		pack[randomIndex].CardValue = 0
		return AddCardValue
	}

}
