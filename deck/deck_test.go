package deck

import (
	"math/big"
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

	for _, suit := range suits {
		if suitsCount[suit] != 13 {
			t.Errorf("Expected 13 cards for suit %s, got %d", suit, suitsCount[suit])
		}
	}

	for _, rank := range ranks {
		if ranksCount[rank] != 4 {
			t.Errorf("Expected 4 cards for rank %s, got %d", rank, ranksCount[rank])
		}
	}
}

func TestNChooseK(t *testing.T) {
	tests := []struct {
		n, k     int
		expected *big.Int
	}{
		{5, 2, big.NewInt(10)},
		{6, 3, big.NewInt(20)},
		{10, 5, big.NewInt(252)},
		{0, 0, big.NewInt(1)},
		{10, 0, big.NewInt(1)},
		{10, 10, big.NewInt(1)},
		{10, 11, big.NewInt(0)},
		{52, 5, big.NewInt(2598960)},
		{52, 8, big.NewInt(752538150)},
	}

	for _, test := range tests {
		result := NChooseK(test.n, test.k)
		if result.Cmp(test.expected) != 0 {
			t.Errorf("NChooseK(%d, %d) = %d; expected %d", test.n, test.k, result, test.expected)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected *big.Int
	}{
		{0, big.NewInt(1)},
		{1, big.NewInt(1)},
		{2, big.NewInt(2)},
		{3, big.NewInt(6)},
		{4, big.NewInt(24)},
		{5, big.NewInt(120)},
		{6, big.NewInt(720)},
		{7, big.NewInt(5040)},
	}

	for _, test := range tests {
		result := Factorial(test.n)
		if result.Cmp(test.expected) != 0 {
			t.Errorf("Factorial(%d) = %d; expected %d", test.n, result, test.expected)
		}
	}
}

func TestProbabilityStraightFlush(t *testing.T) {
	deck := NewStandardDeck()
	expected := 40.0 / 2598960.0
	tests := []struct {
		deck     *Deck
		handSize int
		expected float64
	}{
		{deck, 5, 40.0 / 2598960.0},
		{deck, 8, .00014},
	}

	for _, test := range tests {
		result := CalculateStraightFlushProbability(test.deck, test.handSize)
		if result != expected {
			t.Errorf("ProbabilityStraightFlush() = %f; expected %f", result, test.expected)
		}
	}
}

func TestCountStraightFlushes(t *testing.T) {
	deck := NewStandardDeck()
	count := CountStraightFlushes(deck)
	expected := big.NewInt(40) // There are 40 possible straight flushes in a standard deckb
	if count.Cmp(expected) != 0 {
		t.Errorf("CountStraightFlushes() = %d; expected %d", count, expected)
	}
}

func TestContainsStraightFlush(t *testing.T) {
	hand := []Card{
		{"Hearts", "A"},
		{"Hearts", "2"},
		{"Hearts", "3"},
		{"Hearts", "4"},
		{"Hearts", "5"},
	}
	if !ContainsStraightFlush(hand) {
		t.Errorf("Expected hand to contain a straight flush")
	}
	handRegular := []Card{
		{"Spades", "8" },
		{"Spades", "9"},
		{"Spades", "10"},
		{"Spades", "J"},
		{"Spades", "Q"},
	}
	if !ContainsStraightFlush(handRegular) {
		t.Errorf("Expected hand to contain a straight flush (8-9-10-J-Q)")
	}

	handNoStraightFlush := []Card{
		{ "Clubs", "2"},
		{ "Hearts", "3"},
		{ "Diamonds", "4"},
		{ "Spades", "5"},
		{ "Clubs", "6"},
	}
	if ContainsStraightFlush(handNoStraightFlush) {
		t.Errorf("Expected hand to NOT contain a straight flush")
	}
}

func TestGenerateHands(t *testing.T) {
	deck := NewStandardDeck()
	hands := GenerateHands(deck, 5)
	if hands != 40 { // 52 choose 5
		t.Errorf("Expected %d hands, got %d", 40, hands)
	}
}
