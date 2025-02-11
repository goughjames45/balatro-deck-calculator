package main

import (
	"fmt"
	deckBuilder "github.com/goughjames45/balatro-deck-calculator/internal/deck"
	simulation "github.com/goughjames45/balatro-deck-calculator/internal/probability-simulation"
)

func main() {
	// deck := deckBuilder.NewStandardDeck()
	deck := deckBuilder.NewStandardDeck()
	fmt.Printf("%v\n", len(deck.Cards))
	simulator := simulation.NewProbabilitySimulation(deck)
	fmt.Printf("%+v\n", *simulator.SimulatePokerHandProbabilities(5, 100000))
	fmt.Printf("%+v\n", *simulator.SimulatePokerHandProbabilities(8, 100000))

	removeCards := []deckBuilder.Card{
		deckBuilder.Card{Suit: "Hearts", Rank: "2"},
		deckBuilder.Card{Suit: "Hearts", Rank: "3"},
		deckBuilder.Card{Suit: "Hearts", Rank: "4"},
		deckBuilder.Card{Suit: "Hearts", Rank: "5"},
		deckBuilder.Card{Suit: "Hearts", Rank: "6"},
		deckBuilder.Card{Suit: "Hearts", Rank: "7"},
		deckBuilder.Card{Suit: "Hearts", Rank: "8"},
		deckBuilder.Card{Suit: "Hearts", Rank: "9"},
		deckBuilder.Card{Suit: "Hearts", Rank: "10"},
		deckBuilder.Card{Suit: "Hearts", Rank: "J"},
		deckBuilder.Card{Suit: "Hearts", Rank: "Q"},
		deckBuilder.Card{Suit: "Hearts", Rank: "K"},
		deckBuilder.Card{Suit: "Hearts", Rank: "A"},
		deckBuilder.Card{Suit: "Spades", Rank: "2"},
		deckBuilder.Card{Suit: "Spades", Rank: "3"},
		deckBuilder.Card{Suit: "Spades", Rank: "4"},
		deckBuilder.Card{Suit: "Spades", Rank: "5"},
		deckBuilder.Card{Suit: "Spades", Rank: "6"},
		deckBuilder.Card{Suit: "Spades", Rank: "7"},
		deckBuilder.Card{Suit: "Spades", Rank: "8"},
		deckBuilder.Card{Suit: "Spades", Rank: "9"},
		deckBuilder.Card{Suit: "Spades", Rank: "10"},
		deckBuilder.Card{Suit: "Spades", Rank: "J"},
		deckBuilder.Card{Suit: "Spades", Rank: "Q"},
		deckBuilder.Card{Suit: "Spades", Rank: "K"},
		deckBuilder.Card{Suit: "Spades", Rank: "A"},
		deckBuilder.Card{Suit: "Diamonds", Rank: "2"},
		deckBuilder.Card{Suit: "Diamonds", Rank: "3"},
		deckBuilder.Card{Suit: "Diamonds", Rank: "4"},
		deckBuilder.Card{Suit: "Diamonds", Rank: "5"},
		deckBuilder.Card{Suit: "Diamonds", Rank: "6"},
		deckBuilder.Card{Suit: "Diamonds", Rank: "7"},
		deckBuilder.Card{Suit: "Diamonds", Rank: "8"},
		deckBuilder.Card{Suit: "Diamonds", Rank: "9"},
		deckBuilder.Card{Suit: "Diamonds", Rank: "10"},
		deckBuilder.Card{Suit: "Diamonds", Rank: "J"},
		deckBuilder.Card{Suit: "Diamonds", Rank: "Q"},
		deckBuilder.Card{Suit: "Diamonds", Rank: "K"},
		deckBuilder.Card{Suit: "Diamonds", Rank: "A"},
	}
	newDeck := deckBuilder.NewStandardDeck(deckBuilder.WithTheHangedMan(removeCards))

	fmt.Printf("%+v\n", newDeck.Cards)
	fmt.Printf("%v\n", len(newDeck.Cards))
	newSimulator := simulation.NewProbabilitySimulation(newDeck)
	fmt.Printf("%+v\n", *newSimulator.SimulatePokerHandProbabilities(8, 100000))
}
