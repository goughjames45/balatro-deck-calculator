package deck

import (
	"math/rand"
	"time"
)

type BalatroDeck interface {
	Shuffle()
	DealHand(handSize int) []Card
}

type Deck struct {
	Cards []Card
}

// Card represents a playing card with a suit and a rank.
type Card struct {
	Suit string
	Rank string
}

var Suits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
var Ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

// NewDeck creates and returns a new deck of 52 playing cards.
func NewStandardDeck() *Deck {
	var deck []Card
	for _, suit := range Suits {
		for _, rank := range Ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}
	return &Deck{Cards: deck}
}

func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d Deck) DealHand(handSize int) []Card {
	var hand []Card

	for i := 0; i < handSize; i++ {
		hand = append(hand, d.Cards[i])
	}

	return hand

}
