package hand

import (
	"../deck"
	"math/rand"
	"testing"
)

func TestDeck(t *testing.T) {
	rand.Seed(1)

	t.Run("Hand value without aces calculated correctly", func(t *testing.T) {
		hand1 := Hand{deck.Card{2, 'H'}, deck.Card{2, 'D'}}
		if hand1.Value() != 4 {
			t.Errorf("Expected hand value to be 4, got %d", hand1.Value())
		}

		handWithFaceCard := Hand{deck.Card{'Q', 'H'}, deck.Card{2, 'D'}}
		if handWithFaceCard.Value() != 12 {
			t.Errorf("Expected hand value to be 12, got %d", handWithFaceCard.Value())
		}
	})

	t.Run("Hand value with aces calculated correctly", func(t *testing.T) {
		handWithAceHigh := Hand{deck.Card{'A', 'H'}, deck.Card{3, 'D'}}
		if handWithAceHigh.Value() != 14 {
			t.Errorf("Expected hand value to be 14, got %d", handWithAceHigh.Value())
		}

		handWithAceLow := Hand{deck.Card{'A', 'H'}, deck.Card{'Q', 'D'}, deck.Card{4, 'D'}}
		if handWithAceLow.Value() != 15 {
			t.Errorf("Expected hand value to be 15, got %d", handWithAceLow.Value())
		}

		handWithTwoAces := Hand{deck.Card{'A', 'H'}, deck.Card{'A', 'C'}}
		if handWithTwoAces.Value() != 12 {
			t.Errorf("Expected hand value to be 12, got %d", handWithTwoAces.Value())
		}

		handWithThreeAces := Hand{deck.Card{'A', 'H'}, deck.Card{'A', 'C'}, deck.Card{'A', 'D'}}
		if handWithThreeAces.Value() != 13 {
			t.Errorf("Expected hand value to be 13, got %d", handWithThreeAces.Value())
		}

	})

	t.Run("DealHand returns hand", func(t *testing.T) {
		shuffledDeck := deck.ShuffledDeck()
		expectedHand := Hand{deck.Card{Rank: 10, Suit: 'H'}, deck.Card{Rank: 7, Suit: 'H'}, deck.Card{Rank: 'J', Suit: 'S'}, deck.Card{Rank: 'Q', Suit: 'S'}}
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

	t.Run("StringRepresentation gives correct player representation", func(t *testing.T) {
		hand := Hand{deck.Card{10, 'H'}, deck.Card{'Q', 'C'}}
		expectedRepr := "10 of H, Q of C"
		repr := StringRepresentation(hand, true)
		if repr != expectedRepr {
			t.Errorf("Expected %v got %v", expectedRepr, repr)
		}
	})

	t.Run("StringRepresentation gives correct dealer representation", func(t *testing.T) {
		hand := Hand{deck.Card{10, 'H'}, deck.Card{'Q', 'C'}}
		expectedRepr := "10 of H, [hidden card]"
		repr := StringRepresentation(hand, false)
		if repr != expectedRepr {
			t.Errorf("Expected %v got %v", expectedRepr, repr)
		}
	})
}
