package probabilitySimulation

import (
	d "github.com/goughjames45/balatro-deck-calculator/internal/deck"
)

type ProbabilitySimulation struct {
	deck d.BalatroDeck
}


func NewProbabilitySimulation(bDeck d.BalatroDeck) *ProbabilitySimulation {
	return &ProbabilitySimulation { deck: bDeck }
}


func (p ProbabilitySimulation) SimulatePokerHandProbabilities(handSize, numSimulations int) float64 {
	n := numSimulations
	count := 0

	for i := 0; i <= n; i++ {
		p.deck.Shuffle()
		hand := p.deck.DealHand(handSize)
		
		if ContainsStraightFlush(hand) {
			count++
		}
	}

	return float64(count) / float64(n)

}

// func CalculateStraightFlushProbability(deck d.BalatroDeck, handSize int) float64 {
//
// 	totalHands := new(big.Float).SetInt(NChooseK(len(deck.Cards), handSize))
// 	totalStraightFlushes := float64(GenerateHands(deck, handSize))
// 	fmt.Printf("%v \n", totalStraightFlushes)
//
// 	fHands, _ := totalHands.Float64()
// 	fmt.Printf("%v \n", fHands)
//
// 	return totalStraightFlushes / fHands
// }
//
// // NChooseK calculates the binomial coefficient (n choose k)
// func NChooseK(n, k int) *big.Int {
// 	if k > n {
// 		return big.NewInt(int64(0))
// 	}
//
// 	return Factorial(n).Div(Factorial(n), Factorial(k).Mul(Factorial(k), Factorial(n-k)))
// }
//
// func Factorial(n int) *big.Int {
// 	if n == 0 || n == 1 {
// 		return big.NewInt(int64(1))
// 	}
// 	result := big.NewInt(1)
// 	for i := 2; i <= n; i++ {
// 		x := big.NewInt(int64(i))
// 		result.Mul(result, x)
// 	}
// 	return result
// }
//
// func GenerateHands(deck d.BalatroDeck, handSize int) int {
// 	count := 0
// 	var generate func(start int, hand []Card)
// 	generate = func(start int, hand []Card) {
// 		if len(hand) == handSize {
// 			if ContainsStraightFlush(hand) {
// 				count += 1
// 			}
// 			return
// 		}
// 		for i := start; i < len(deck.Cards); i++ {
// 			generate(i+1, append(hand, deck.Cards[i]))
// 		}
// 	}
// 	generate(0, []Card{})
// 	return int(count)
// }

func ContainsStraightFlush(hand []d.Card) bool {
	for _, suit := range d.Suits {
		for i := 0; i <= 8; i++ {
			hasStraightFlush := true
			for j := 0; j < 5; j++ {
				card := d.Card{Rank: d.Ranks[i+j], Suit: suit}
				found := false
				for _, c := range hand {
					if c == card {
						found = true
						break
					}
				}
				if !found {
					hasStraightFlush = false
					break
				}
			}
			if hasStraightFlush {
				return true
			}
		}
		// Check for A-2-3-4-5 straight flush
		aceLowStraight := []string{"A", "2", "3", "4", "5"}
		hasAceLowStraightFlush := true
		for _, rank := range aceLowStraight {
			card := d.Card{Rank: rank, Suit: suit}
			found := false
			for _, c := range hand {
				if c == card {
					found = true
					break
				}
			}
			if !found {
				hasAceLowStraightFlush = false
				break
			}
		}
		if hasAceLowStraightFlush {
			return true
		}
	}
	return false
}

// func CountStraightFlushes(deck d.BalatroDeck) *big.Int {
// 	count := 0
// 	for _, suit := range suits {
// 		for i := 0; i <= 8; i++ {
// 			hasStraightFlush := true
// 			for j := 0; j < 5; j++ {
// 				card := Card{Rank: ranks[i+j], Suit: suit}
// 				found := false
// 				for _, c := range deck.Cards {
// 					if c == card {
// 						found = true
// 						break
// 					}
// 				}
// 				if !found {
// 					hasStraightFlush = false
// 					break
// 				}
// 			}
// 			if hasStraightFlush {
// 				count++
// 			}
// 		}
// 		// Check for A-2-3-4-5 straight flush
// 		aceLowStraight := []string{"A", "2", "3", "4", "5"}
// 		hasAceLowStraightFlush := true
// 		for _, rank := range aceLowStraight {
// 			card := Card{Rank: rank, Suit: suit}
// 			found := false
// 			for _, c := range deck.Cards {
// 				if c == card {
// 					found = true
// 					break
// 				}
// 			}
// 			if !found {
// 				hasAceLowStraightFlush = false
// 				break
// 			}
// 		}
// 		if hasAceLowStraightFlush {
// 			count++
// 		}
// 	}
// 	return big.NewInt(int64(count))
// }
