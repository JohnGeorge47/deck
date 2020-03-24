package blackjack

import (
	"github.com/JohnGeorge47/deck/card"
)

type Player struct {
	PlayerId    int
	PlayerCards []card.Card
	HiddenCard  card.Card
	ShownCard   []card.Card
	PlayerScore int
}

func NewPlayers() []Player {
	playersArray := make([]Player, 6,6)
	for i := 0; i < 6; i++ {
		p:=Player{
			PlayerId:    i,
		}
		playersArray[i] = p
	}
	return playersArray
}
