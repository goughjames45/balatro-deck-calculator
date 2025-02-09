package main

import (
	"fmt"
	deckBuilder "github.com/goughjames45/balatro-deck-calculator/internal/deck"
	simulation "github.com/goughjames45/balatro-deck-calculator/internal/probability-simulation"
)

func main() {
	// deck := deckBuilder.NewStandardDeck()
	deck := deckBuilder.NewStandardDeck()
	simulator := simulation.NewProbabilitySimulation(deck)
	fmt.Printf("%+v\n", simulator.SimulatePokerHandProbabilities(5, 100000))
	fmt.Printf("%+v\n", simulator.SimulatePokerHandProbabilities(8,  100000))
}
