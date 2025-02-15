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


func TestFuncOps(t *testing.T) {
	tests := []struct {
		name string
		op FuncOp
		expectedDeckSize int
		expectedSuits map[string]int
		expectedRanks map[string]int
	}{
		{
			name: "With the Hanged Man",
			op: WithTheHangedMan(
				[]Card{
					{Suit: "Hearts", Rank: "2"},
				},
			),
			expectedDeckSize: 51,
			expectedSuits: map[string]int {
				"Hearts": 12,
				"Spades": 13,
				"Diamonds": 13,
				"Clubs": 13,
			},
			expectedRanks: map[string]int {
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
			},
		},
		{
			name: "With The Sun",
			op: WithTheSun(
				[]Card{
					{Suit: "Clubs", Rank: "2"},
					{Suit: "Clubs", Rank: "3"},
					{Suit: "Clubs", Rank: "4"},
				},
			),
			expectedDeckSize: 52,
			expectedSuits: map[string]int {
				"Hearts": 16,
				"Spades": 13,
				"Diamonds": 13,
				"Clubs": 10,
			},
			expectedRanks: map[string]int {
				"2": 4,
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
			},
		},
		{
			name: "With The Star",
			op: WithTheStar(
				[]Card{
					{Suit: "Clubs", Rank: "2"},
					{Suit: "Clubs", Rank: "3"},
					{Suit: "Clubs", Rank: "4"},
				},
			),
			expectedDeckSize: 52,
			expectedSuits: map[string]int {
				"Hearts": 13,
				"Spades": 13,
				"Diamonds": 16,
				"Clubs": 10,
			},
			expectedRanks: map[string]int {
				"2": 4,
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
			},
		},
		{
			name: "With The Moon",
			op: WithTheMoon(
				[]Card{
					{Suit: "Hearts", Rank: "2"},
					{Suit: "Hearts", Rank: "3"},
					{Suit: "Hearts", Rank: "4"},
				},
			),
			expectedDeckSize: 52,
			expectedSuits: map[string]int {
				"Hearts": 10,
				"Spades": 13,
				"Diamonds": 13,
				"Clubs": 16,
			},
			expectedRanks: map[string]int {
				"2": 4,
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
			},
		},
		{
			name: "With The World",
			op: WithTheWorld(
				[]Card{
					{Suit: "Hearts", Rank: "2"},
					{Suit: "Hearts", Rank: "3"},
					{Suit: "Hearts", Rank: "4"},
				},
			),
			expectedDeckSize: 52,
			expectedSuits: map[string]int {
				"Hearts": 10,
				"Spades": 16,
				"Diamonds": 13,
				"Clubs": 13,
			},
			expectedRanks: map[string]int {
				"2": 4,
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
			},
		},
		{
			name: "With Strength",
			op: WithStrength(
				[]Card{
					{Suit: "Hearts", Rank: "3"},
					{Suit: "Spades", Rank: "3"},
				},
			),
			expectedDeckSize: 52,
			expectedSuits: map[string]int {
				"Hearts": 13,
				"Spades": 13,
				"Diamonds": 13,
				"Clubs": 13,
			},
			expectedRanks: map[string]int {
				"2": 4,
				"3": 2,
				"4": 6,
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
			},
		},
		{
			name: "With Death",
			op: WithDeath(
				[]Card{
					{Suit: "Hearts", Rank: "3"},
					{Suit: "Spades", Rank: "5"},
				},
			),
			expectedDeckSize: 52,
			expectedSuits: map[string]int {
				"Hearts": 12,
				"Spades": 14,
				"Diamonds": 13,
				"Clubs": 13,
			},
			expectedRanks: map[string]int {
				"2": 4,
				"3": 3,
				"4": 4,
				"5": 5,
				"6": 4,
				"7": 4,
				"8": 4,
				"9": 4,
				"10": 4,
				"J": 4,
				"Q": 4,
				"K": 4,
				"A": 4,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			suitsCount := make(map[string]int)
			ranksCount := make(map[string]int)

			deck := NewStandardDeck(test.op)

			for _, card := range deck.Cards {
				suitsCount[card.Suit]++
				ranksCount[card.Rank]++
			}

			if len(deck.Cards) != test.expectedDeckSize {
				t.Errorf("Expected deck length of %d, got %d", test.expectedDeckSize, len(deck.Cards))
			}

			for _, suit := range Suits {
				if suitsCount[suit] != test.expectedSuits[suit] {
					t.Errorf("Expected %d cards for suit %s, got %d", test.expectedSuits[suit], suit, suitsCount[suit])
				}
			}

			for _, rank := range Ranks {
				if ranksCount[rank] != test.expectedRanks[rank] {
					t.Errorf("Expected %d cards for rank %s, got %d", test.expectedRanks[rank], rank, ranksCount[rank])
				}
			}
		})
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
