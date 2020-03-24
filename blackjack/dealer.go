package blackjack

import (
	"github.com/JohnGeorge47/deck/card"
)

type Dealer struct {
	DealerCards []card.Card
	ShownCard   []card.Card
	HiddenCard  card.Card
}

func NewDealer() Dealer {
	dealer := Dealer{
		DealerCards: nil,
		ShownCard:   nil,
		HiddenCard:  card.Card{},
	}
	return dealer
}
