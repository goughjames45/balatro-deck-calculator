package main

import (
	"fmt"
	deckBuilder "github.com/goughjames45/balatro-deck-calculator/deck"
)

func main() {
	// deck := deckBuilder.NewStandardDeck()
	fmt.Printf("%+v\n", deckBuilder.SimulateStraightFlushProbability(5))
	fmt.Printf("%+v\n", deckBuilder.SimulateStraightFlushProbability(8))
}
