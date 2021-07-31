package hand

import (
	"../deck"
	"math/rand"
	"testing"
)

func TestDeck(t *testing.T) {
	rand.Seed(1)

	t.Run("DealHand returns hand", func(t *testing.T) {
		shuffledDeck := deck.ShuffledDeck()
		expectedHand := Hand{deck.Card{Value: 10, Suit: 'H'}, deck.Card{Value: 7, Suit: 'H'}, deck.Card{Value: 'J', Suit: 'S'}, deck.Card{Value: 'Q', Suit: 'S'}}
		hand := DealHand(&shuffledDeck)

		if len(hand) != 2 {
			t.Errorf("hand length is %v", len(hand))
		}
		for i, card := range hand {
			if card != expectedHand[i] {
				t.Errorf("got %v want %v", card, expectedHand[i])
			}
		}
	})

	t.Run("DealHand removes cards from deck", func(t *testing.T) {
		shuffledDeck := deck.ShuffledDeck()
		DealHand(&shuffledDeck)
		DealHand(&shuffledDeck)
		if len(shuffledDeck) != 48 {
			t.Errorf("Expected deck size of 50, got %v", len(shuffledDeck))
		}
	})

	t.Run("StringRepresentation gives correct representation", func(t *testing.T) {
		hand := Hand{deck.Card{10, 'H'}, deck.Card{'Q', 'C'}}
		expectedRepr := "10 of H, Q of C"
		repr := StringRepresentation(hand)
		if repr != expectedRepr {
			t.Errorf("Expected %v got %v", expectedRepr, repr)
		}
	})
}
