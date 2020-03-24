package main

import (
	"fmt"
	"github.com/JohnGeorge47/deck/blackjack"
	"github.com/JohnGeorge47/deck/card"
)

const (
	dealer=0
)
type BlackJack struct {
	players     []blackjack.Player
	currentDeck []card.Card
}

func (b *BlackJack) DealCards(){
	for i := 0; i < len(b.players); i++ {
		b.players[i].PlayerCards = append(b.players[i].PlayerCards, b.currentDeck[0])
		b.currentDeck = b.currentDeck[1:]
	}
}

func NewBlackJackGame() *BlackJack {
	return &BlackJack{
		players:     blackjack.NewPlayers(),
		currentDeck: card.New(card.MultipleDecks(3), card.Shuffle),
	}
}

func main() {
	fmt.Println("Welcome to this game of blackjack")
	fmt.Println("We are open to deal cards now")
	game := NewBlackJackGame()
	for i:=0;i<2;i++{
		game.DealCards()
		for i := 0; i < len(game.players); i++ {
			if i==dealer&&i==0{
				fmt.Println(game.players[i].PlayerCards,game.players[i].PlayerId)
			}
		}
	}
}
