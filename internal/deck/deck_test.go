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


	removeCards := []Card{
		Card{Suit: "Hearts", Rank: "2"},
	}
	deck = NewStandardDeck(WithTheHangedMan(removeCards))
	if len(deck.Cards) != 51 {
		t.Errorf("Expected deck length of 52, got %d", len(deck.Cards))
	}
	clear(suitsCount)
	clear(ranksCount)
	expectedSuitsCount := map[string]int {
		"Hearts": 12,
		"Spades": 13,
		"Diamonds": 13,
		"Clubs": 13,
	}
	expectedRanksCount := map[string]int {
		"2": 3,
		"3": 4,
		"4": 4,
		"5": 4,
		"6": 4,
		"7": 4,
		"8": 4,
		"9": 4,
		"10": 4,
		"J": 4,
		"Q": 4,
		"K": 4,
		"A": 4,
	}
	for _, card := range deck.Cards {
		suitsCount[card.Suit]++
		ranksCount[card.Rank]++
	}
	for _, suit := range Suits {
		if suitsCount[suit] != expectedSuitsCount[suit] {
			t.Errorf("Expected 13 cards for suit %s, got %d", suit, suitsCount[suit])
		}
	}

	for _, rank := range Ranks {
		if ranksCount[rank] != expectedRanksCount[rank] {
			t.Errorf("Expected 4 cards for rank %s, got %d", rank, ranksCount[rank])
		}
	}
}

func TestDealHand(t *testing.T) {

	deck := NewStandardDeck()
	hand := deck.DealHand(8)

	if len(hand) != 8 {
		t.Errorf("Expected hand of size %d, got %d", 8, len(hand))
	}
}

func TestShuffle(t *testing.T) {
	deck := NewStandardDeck()
	originalDeck := append([]Card{}, deck.Cards...)

	deck.Shuffle()
	different := false
	for i, card := range deck.Cards {
		if card != originalDeck[i] {
			different = true
			break
		}
	}

	if !different {
		t.Errorf("Deck was not shuffled")
	}
}
