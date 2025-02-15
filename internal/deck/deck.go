package deck

import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
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

type FuncOp func(*Deck)

var Suits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
var Ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

// NewDeck creates and returns a new deck of 52 playing cards.
func NewStandardDeck(opts ...FuncOp) *Deck {
	var deck []Card
	for _, suit := range Suits {
		for _, rank := range Ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}

	dck := &Deck{Cards: deck}

	for _, opt := range opts {
		opt(dck)
	}

	return dck
}

func WithTheHangedMan(card []Card) func(*Deck){
	return func(d *Deck) {
		for _, c := range card {
			for i, other := range d.Cards {
				if c == other {
				    d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
				}
			}
		}
	}
}

func WithStrength(card []Card) func(*Deck){
	return func(d *Deck) {
		for _, c := range card {
			for i, other := range d.Cards {
				if c == other {
					rank, err := strconv.Atoi(c.Rank)
					if err != nil {
						fmt.Println("Error:", err)
						return
					}

					d.Cards[i] = Card{Suit: c.Suit, Rank: strconv.Itoa(rank+1)}
				}
			}
		}
	}
}

func WithDeath(card []Card) func(*Deck){
	return func(d *Deck) {
		for i, other := range d.Cards {
			if card[0] == other {
				d.Cards[i] = Card{Suit: card[1].Suit, Rank: card[1].Rank}
			}
		}
	}
}

func WithTheStar(card []Card) func(*Deck){
	return func(d *Deck) {
		switchCardsToSuit(card, d, "Diamonds")
	}
}

func WithTheMoon(card []Card) func(*Deck){
	return func(d *Deck) {
		switchCardsToSuit(card, d, "Clubs")
	}
}

func WithTheSun(card []Card) func(*Deck){
	return func(d *Deck) {
		switchCardsToSuit(card, d, "Hearts")
	}
}

func WithTheWorld(card []Card) func(*Deck){
	return func(d *Deck) {
		switchCardsToSuit(card, d, "Spades")
	}
}

func switchCardsToSuit(card []Card, d *Deck, suit string) {
		for _, c := range card {
			for i, other := range d.Cards {
				if c == other {
					d.Cards[i] = Card{Suit: suit, Rank: c.Rank}
				}
			}
		}
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
