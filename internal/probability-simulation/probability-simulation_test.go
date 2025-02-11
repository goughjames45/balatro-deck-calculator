package probabilitySimulation

import (
	d "github.com/goughjames45/balatro-deck-calculator/internal/deck"
	"testing"
)

// func TestNChooseK(t *testing.T) {
// 	tests := []struct {
// 		n, k     int
// 		expected *big.Int
// 	}{
// 		{5, 2, big.NewInt(10)},
// 		{6, 3, big.NewInt(20)},
// 		{10, 5, big.NewInt(252)},
// 		{0, 0, big.NewInt(1)},
// 		{10, 0, big.NewInt(1)},
// 		{10, 10, big.NewInt(1)},
// 		{10, 11, big.NewInt(0)},
// 		{52, 5, big.NewInt(2598960)},
// 		{52, 8, big.NewInt(752538150)},
// 	}
//
// 	for _, test := range tests {
// 		result := NChooseK(test.n, test.k)
// 		if result.Cmp(test.expected) != 0 {
// 			t.Errorf("NChooseK(%d, %d) = %d; expected %d", test.n, test.k, result, test.expected)
// 		}
// 	}
// }
//
// func TestFactorial(t *testing.T) {
// 	tests := []struct {
// 		n        int
// 		expected *big.Int
// 	}{
// 		{0, big.NewInt(1)},
// 		{1, big.NewInt(1)},
// 		{2, big.NewInt(2)},
// 		{3, big.NewInt(6)},
// 		{4, big.NewInt(24)},
// 		{5, big.NewInt(120)},
// 		{6, big.NewInt(720)},
// 		{7, big.NewInt(5040)},
// 	}
//
// 	for _, test := range tests {
// 		result := Factorial(test.n)
// 		if result.Cmp(test.expected) != 0 {
// 			t.Errorf("Factorial(%d) = %d; expected %d", test.n, result, test.expected)
// 		}
// 	}
// }

// func TestProbabilityStraightFlush(t *testing.T) {
// 	deck := NewStandardDeck()
// 	expected := 40.0 / 2598960.0
// 	tests := []struct {
// 		deck     *Deck
// 		handSize int
// 		expected float64
// 	}{
// 		{deck, 5, 40.0 / 2598960.0},
// 		{deck, 8, .00014},
// 	}
//
// 	for _, test := range tests {
// 		result := CalculateStraightFlushProbability(test.deck, test.handSize)
// 		if result != expected {
// 			t.Errorf("ProbabilityStraightFlush() = %f; expected %f", result, test.expected)
// 		}
// 	}
// }

// func TestCountStraightFlushes(t *testing.T) {
// 	deck := NewStandardDeck()
// 	count := CountStraightFlushes(deck)
// 	expected := big.NewInt(40) // There are 40 possible straight flushes in a standard deckb
// 	if count.Cmp(expected) != 0 {
// 		t.Errorf("CountStraightFlushes() = %d; expected %d", count, expected)
// 	}
// }

func TestContainsStraightFlush(t *testing.T) {
	hand := []d.Card{
		{"Hearts", "A"},
		{"Hearts", "2"},
		{"Hearts", "3"},
		{"Hearts", "4"},
		{"Hearts", "5"},
	}
	if !ContainsStraightFlush(hand) {
		t.Errorf("Expected hand to contain a straight flush")
	}
	handRegular := []d.Card{
		{"Spades", "8"},
		{"Spades", "9"},
		{"Spades", "10"},
		{"Spades", "J"},
		{"Spades", "Q"},
	}
	if !ContainsStraightFlush(handRegular) {
		t.Errorf("Expected hand to contain a straight flush (8-9-10-J-Q)")
	}

	handNoStraightFlush := []d.Card{
		{"Clubs", "2"},
		{"Hearts", "3"},
		{"Diamonds", "4"},
		{"Spades", "5"},
		{"Clubs", "6"},
	}
	if ContainsStraightFlush(handNoStraightFlush) {
		t.Errorf("Expected hand to NOT contain a straight flush")
	}
}

func TestContainsFlush(t *testing.T) {
	hand := []d.Card{
		{"Hearts", "A"},
		{"Hearts", "2"},
		{"Hearts", "3"},
		{"Hearts", "4"},
		{"Hearts", "5"},
	}
	if !ContainsFlush(hand) {
		t.Errorf("Expected hand to contain a flush")
	}

	handNoFlush := []d.Card{
		{"Clubs", "2"},
		{"Hearts", "3"},
		{"Diamonds", "4"},
		{"Spades", "5"},
		{"Clubs", "6"},
	}
	if ContainsFlush(handNoFlush) {
		t.Errorf("Expected hand to NOT contain a flush")
	}
}

func TestContainsStraight(t *testing.T) {
	handWithStraight := []d.Card{
		{"Hearts", "6"},
		{"Diamonds", "7"},
		{"Clubs", "8"},
		{"Spades", "9"},
		{"Hearts", "10"},
	}
	if !ContainsStraight(handWithStraight) {
		t.Errorf("Expected hand to contain a straight")
	}

	handWithoutStraight := []d.Card{
		{"Hearts", "2"},
		{"Diamonds", "3"},
		{"Clubs", "5"},
		{"Spades", "7"},
		{"Hearts", "9"},
	}
	if ContainsStraight(handWithoutStraight) {
		t.Errorf("Expected hand to NOT contain a straight")
	}
	handWithStraight = []d.Card{
		{"Hearts", "A"},
		{"Hearts", "2"},
		{"Diamonds", "3"},
		{"Clubs", "4"},
		{"Spades", "5"},
	}
	if !ContainsStraight(handWithStraight) {
		t.Errorf("Expected hand to contain a straight")
	}
	handWithStraight = []d.Card{
		{"Hearts", "A"},
		{"Hearts", "2"},
		{"Diamonds", "3"},
		{"Clubs", "4"},
		{"Spades", "5"},
		{"Spades", "9"},
		{"Spades", "Q"},
	}
	if !ContainsStraight(handWithStraight) {
		t.Errorf("Expected hand to contain a straight")
	}
}

func TestContainsFourOfAKind(t *testing.T) {
	handWithFour := []d.Card{
		{"Hearts", "K"},
		{"Diamonds", "K"},
		{"Clubs", "K"},
		{"Spades", "K"},
		{"Hearts", "7"},
	}
	if !ContainsFourOfAKind(handWithFour) {
		t.Errorf("Expected hand to contain four of a kind")
	}

	handWithoutFour := []d.Card{
		{"Hearts", "2"},
		{"Diamonds", "3"},
		{"Clubs", "4"},
		{"Spades", "5"},
		{"Hearts", "6"},
	}
	if ContainsFourOfAKind(handWithoutFour) {
		t.Errorf("Expected hand to NOT contain four of a kind")
	}
}

func TestContainsFullHouse(t *testing.T) {
	handWithFullHouse := []d.Card{
		{"Hearts", "Q"},
		{"Diamonds", "Q"},
		{"Clubs", "Q"},
		{"Spades", "8"},
		{"Hearts", "8"},
	}
	if !ContainsFullHouse(handWithFullHouse) {
		t.Errorf("Expected hand to contain a full house")
	}

	handWithoutFullHouse := []d.Card{
		{"Hearts", "2"},
		{"Diamonds", "3"},
		{"Clubs", "4"},
		{"Spades", "5"},
		{"Hearts", "6"},
	}
	if ContainsFullHouse(handWithoutFullHouse) {
		t.Errorf("Expected hand to NOT contain a full house")
	}
}

func TestContainsThreeOfAKind(t *testing.T) {
	handWithThree := []d.Card{
		{"Hearts", "K"},
		{"Diamonds", "K"},
		{"Clubs", "K"},
		{"Spades", "Q"},
		{"Hearts", "7"},
	}
	if !ContainsThreeOfAKind(handWithThree) {
		t.Errorf("Expected hand to contain three of a kind")
	}

	handWithoutThree := []d.Card{
		{"Hearts", "2"},
		{"Diamonds", "3"},
		{"Clubs", "4"},
		{"Spades", "5"},
		{"Hearts", "6"},
	}
	if ContainsThreeOfAKind(handWithoutThree) {
		t.Errorf("Expected hand to NOT contain three of a kind")
	}
}

func TestContainsTwoPair(t *testing.T) {
	handWithTwoPair := []d.Card{
		{"Hearts", "6"},
		{"Diamonds", "6"},
		{"Clubs", "9"},
		{"Spades", "9"},
		{"Hearts", "10"},
	}
	if !ContainsTwoPair(handWithTwoPair) {
		t.Errorf("Expected hand to contain two pair")
	}

	handWithoutTwoPair := []d.Card{
		{"Hearts", "6"},
		{"Diamonds", "7"},
		{"Clubs", "8"},
		{"Spades", "9"},
		{"Hearts", "10"},
	}
	if ContainsTwoPair(handWithoutTwoPair) {
		t.Errorf("Expected hand to NOT contain two pair")
	}

	handWithThreeOfAKind := []d.Card{
		{"Hearts", "6"},
		{"Diamonds", "6"},
		{"Clubs", "6"},
		{"Spades", "9"},
		{"Hearts", "9"},
	}
	if ContainsTwoPair(handWithThreeOfAKind) {
		t.Errorf("Expected hand to NOT contain exactly two pair")
	}
}

func TestContainsPair(t *testing.T) {
	handWithPair := []d.Card{
		{"Hearts", "K"},
		{"Diamonds", "K"},
		{"Clubs", "K"},
		{"Spades", "Q"},
		{"Hearts", "7"},
	}
	if !ContainsPair(handWithPair) {
		t.Errorf("Expected hand to contain pair")
	}

	handWithoutPair := []d.Card{
		{"Hearts", "2"},
		{"Diamonds", "3"},
		{"Clubs", "4"},
		{"Spades", "5"},
		{"Hearts", "6"},
	}
	if ContainsPair(handWithoutPair) {
		t.Errorf("Expected hand to NOT contain pair")
	}
}
// func TestGenerateHands(t *testing.T) {
// 	deck := NewStandardDeck()
// 	hands := GenerateHands(deck, 5)
// 	if hands != 40 { // 52 choose 4
// 		t.Errorf("Expected %d hands, got %d", 40, hands)
// 	}
// }
