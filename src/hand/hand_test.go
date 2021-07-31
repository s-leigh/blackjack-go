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

		for i, card := range hand {
			if card != expectedHand[i] {
				t.Errorf("got %v want %v", card, expectedHand[i])
			}
		}
	})

	t.Run("DealHand removes cards from deck", func(t *testing.T) {
		shuffledDeck := deck.ShuffledDeck()
		DealHand(&shuffledDeck)
		if len(shuffledDeck) != 50 {
			t.Errorf("Expected deck size of 50, got %v", len(shuffledDeck))
		}

	})
}
