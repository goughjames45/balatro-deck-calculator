package deck

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewStandardDeck()
	if len(deck.Cards) != 52 {
		t.Errorf("Expected deck length of 52, got %d", len(deck.Cards))
	}

	suitsCount := make(map[string]int)
	ranksCount := make(map[string]int)
	for _, card := range deck.Cards {
		suitsCount[card.Suit]++
		ranksCount[card.Rank]++
	}

	for _, suit := range Suits {
		if suitsCount[suit] != 13 {
			t.Errorf("Expected 13 cards for suit %s, got %d", suit, suitsCount[suit])
		}
	}

	for _, rank := range Ranks {
		if ranksCount[rank] != 4 {
			t.Errorf("Expected 4 cards for rank %s, got %d", rank, ranksCount[rank])
		}
	}
}

