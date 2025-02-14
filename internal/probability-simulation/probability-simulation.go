package probabilitySimulation

import (
	d "github.com/goughjames45/balatro-deck-calculator/internal/deck"
	"sort"
)

type ProbabilitySimulation struct {
	deck d.BalatroDeck
}

type HandAccumulator struct {
	StraightFlushCount float64 
	FourOfAKindCount   float64 
	FullHouseCount     float64 
	FlushCount         float64 
	StraightCount      float64 
	ThreeOfAKindCount  float64 
	TwoPairCount       float64 
	PairCount          float64 
}

func NewProbabilitySimulation(bDeck d.BalatroDeck) *ProbabilitySimulation {
	return &ProbabilitySimulation{deck: bDeck}
}

func NewHandAccumulator() *HandAccumulator {
	return &HandAccumulator{
		StraightFlushCount: 0,
		FourOfAKindCount:   0,
		FullHouseCount:     0,
		FlushCount:         0,
		StraightCount:      0,
		ThreeOfAKindCount:  0,
		TwoPairCount:       0,
		PairCount:               0,
	}
}

func (h *HandAccumulator) ConvertToPercentage(n float64) {
	h.StraightFlushCount = h.StraightFlushCount / n * 100
	h.FourOfAKindCount = h.FourOfAKindCount / n * 100
	h.FullHouseCount = h.FullHouseCount / n * 100
	h.FlushCount = h.FlushCount / n * 100
	h.StraightCount = h.StraightCount / n * 100
	h.ThreeOfAKindCount = h.ThreeOfAKindCount / n * 100
	h.TwoPairCount = h.TwoPairCount / n * 100
	h.PairCount = h.PairCount / n * 100
}

func (p *ProbabilitySimulation) SimulatePokerHandProbabilities(handSize, numSimulations int) *HandAccumulator {
	n := numSimulations
	handCount := NewHandAccumulator()

	for i := 0; i <= n; i++ {
		p.deck.Shuffle()
		hand := p.deck.DealHand(handSize)

		if ContainsStraightFlush(hand) {
			handCount.StraightFlushCount++
		}
		if ContainsFourOfAKind(hand) {
			handCount.FourOfAKindCount++
		}
		if ContainsFullHouse(hand) {
			handCount.FullHouseCount++
		}
		if ContainsFlush(hand) {
			handCount.FlushCount++
		}
		if ContainsStraight(hand) {
			handCount.StraightCount++
		}
		if ContainsThreeOfAKind(hand) {
			handCount.ThreeOfAKindCount++
		}
		if ContainsTwoPair(hand) {
			handCount.TwoPairCount++
		}
		if ContainsPair(hand) {
			handCount.PairCount++
		}
	}

	handCount.ConvertToPercentage(float64(n))

	return handCount 
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

func ContainsFourOfAKind(hand []d.Card) bool {
	rankCount := make(map[string]int)
	for _, card := range hand {
		rankCount[card.Rank]++
	}
	for _, count := range rankCount {
		if count >= 4 {
			return true
		}
	}
	return false
}

func ContainsFullHouse(hand []d.Card) bool {
	rankCount := make(map[string]int)
	for _, card := range hand {
		rankCount[card.Rank]++
	}
	hasThree := false
	hasTwo := false
	for _, count := range rankCount {
		if count >= 3 {
			hasThree = true
		} else if count >= 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
}


func ContainsFlush(hand []d.Card) bool {
	suitCount := make(map[string]int)
	for _, card := range hand {
		suitCount[card.Suit]++
	}
	for _, count := range suitCount {
		if count >= 5 {
			return true
		}
	}
	return false
}

func ContainsStraight(hand []d.Card) bool {
	rankOrder := map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 11, "Q": 12, "K": 13, "A": 14}
	uniqueRanks := make(map[int]bool)
	for _, card := range hand {
		uniqueRanks[rankOrder[card.Rank]] = true
	}
	var ranks []int
	for rank := range uniqueRanks {
		ranks = append(ranks, rank)
	}
	sort.Ints(ranks)

	for i := 0; i <= len(ranks)-5; i++ {
		if ranks[i+4] == ranks[i]+4 {
			return true
		}
	}

	// Check for A-2-3-4-5 special case
	if uniqueRanks[14] && uniqueRanks[2] && uniqueRanks[3] && uniqueRanks[4] && uniqueRanks[5] {
		return true
	}

	return false
}

func ContainsThreeOfAKind(hand []d.Card) bool {
	rankCount := make(map[string]int)
	for _, card := range hand {
		rankCount[card.Rank]++
	}
	for _, count := range rankCount {
		if count >= 3 {
			return true
		}
	}
	return false
}

func ContainsTwoPair(hand []d.Card) bool {
	rankCount := make(map[string]int)
	for _, card := range hand {
		rankCount[card.Rank]++
	}

	pairCount := 0
	for _, count := range rankCount {
		if count >= 2 {
			pairCount++
		}
	}

	return pairCount >= 2
}

func ContainsPair(hand []d.Card) bool {
	rankCount := make(map[string]int)
	for _, card := range hand {
		rankCount[card.Rank]++
	}
	for _, count := range rankCount {
		if count >= 2 {
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
