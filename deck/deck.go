package deck

import (
	"math/big"
	"fmt"
)

type Deck struct {
	Cards []Card
}

// Card represents a playing card with a suit and a rank.
type Card struct {
	Suit string
	Rank string
}

var suits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

// NewDeck creates and returns a new deck of 52 playing cards.
func NewStandardDeck() *Deck {
	var deck []Card
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}
	return &Deck{Cards: deck}
}

func CalculateStraightFlushProbability(deck *Deck, handSize int) float64 {

	totalHands := new(big.Float).SetInt(NChooseK(len(deck.Cards), handSize))
	totalStraightFlushes := new(big.Float).SetInt(CountStraightFlushes(deck))

	fHands, _ := totalHands.Float64()
	fmt.Println(fHands)
	fFlushes, _ := totalStraightFlushes.Float64()
	fmt.Println(fFlushes)

	return fFlushes / fHands
}

// NChooseK calculates the binomial coefficient (n choose k)
func NChooseK(n, k int) *big.Int {
	if k > n {
		return big.NewInt(int64(0))
	}

	return Factorial(n).Div(Factorial(n), Factorial(k).Mul(Factorial(k), Factorial(n-k)))
}

func Factorial(n int) *big.Int {
	if n == 0 || n == 1 {
		return big.NewInt(int64(1))
	}
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		x := big.NewInt(int64(i))
		result.Mul(result, x)
	}
	return result
}

func CountStraightFlushes(deck *Deck) *big.Int {
	count := 0
	for _, suit := range suits {
		for i := 0; i <= 8; i++ {
			hasStraightFlush := true
			for j := 0; j < 5; j++ {
				card := Card{Rank: ranks[i+j], Suit: suit}
				found := false
				for _, c := range deck.Cards {
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
				count++
			}
		}
		// Check for A-2-3-4-5 straight flush
		aceLowStraight := []string{"A", "2", "3", "4", "5"}
		hasAceLowStraightFlush := true
		for _, rank := range aceLowStraight {
			card := Card{Rank: rank, Suit: suit}
			found := false
			for _, c := range deck.Cards {
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
			count++
		}
	}
	return big.NewInt(int64(count))
}
