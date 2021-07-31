package hand

import (
	"../deck"
	"math/rand"
	"testing"
)

func TestDeck(t *testing.T) {
	rand.Seed(1)

	t.Run("Hand value without aces calculated correctly", func(t *testing.T) {
		hand1 := Hand{[]deck.Card{{2, 'H'}, {2, 'D'}}, player}
		if hand1.Value() != 4 {
			t.Errorf("Expected hand value to be 4, got %d", hand1.Value())
		}

		handWithFaceCard := Hand{[]deck.Card{{'Q', 'H'}, {2, 'D'}}, player}
		if handWithFaceCard.Value() != 12 {
			t.Errorf("Expected hand value to be 12, got %d", handWithFaceCard.Value())
		}
	})

	t.Run("Hand value with aces calculated correctly", func(t *testing.T) {
		handWithAceHigh := Hand{[]deck.Card{{'A', 'H'}, {3, 'D'}}, player}
		if handWithAceHigh.Value() != 14 {
			t.Errorf("Expected hand value to be 14, got %d", handWithAceHigh.Value())
		}

		handWithAceLow := Hand{[]deck.Card{{'A', 'H'}, {'Q', 'D'}, {4, 'D'}}, player}
		if handWithAceLow.Value() != 15 {
			t.Errorf("Expected hand value to be 15, got %d", handWithAceLow.Value())
		}

		handWithTwoAces := Hand{[]deck.Card{{'A', 'H'}, {'A', 'C'}}, player}
		if handWithTwoAces.Value() != 12 {
			t.Errorf("Expected hand value to be 12, got %d", handWithTwoAces.Value())
		}

		handWithThreeAces := Hand{[]deck.Card{{'A', 'H'}, {'A', 'C'}, {'A', 'D'}}, player}
		if handWithThreeAces.Value() != 13 {
			t.Errorf("Expected hand value to be 13, got %d", handWithThreeAces.Value())
		}

	})

	t.Run("DealHand returns hand", func(t *testing.T) {
		shuffledDeck := deck.ShuffledDeck()
		expectedHand := Hand{[]deck.Card{{Rank: 10, Suit: 'H'}, {Rank: 7, Suit: 'H'}, {Rank: 'J', Suit: 'S'}, {Rank: 'Q', Suit: 'S'}}, player}
		hand := DealHand(&shuffledDeck, false)

		if len(hand.cards) != 2 {
			t.Errorf("hand length is %v", len(hand.cards))
		}
		for i, card := range hand.cards {
			if card != expectedHand.cards[i] {
				t.Errorf("got %v want %v", card, expectedHand.cards[i])
			}
		}
	})

	t.Run("DealHand removes cards from deck", func(t *testing.T) {
		shuffledDeck := deck.ShuffledDeck()
		DealHand(&shuffledDeck, false)
		DealHand(&shuffledDeck, true)
		if len(shuffledDeck) != 48 {
			t.Errorf("Expected deck size of 50, got %v", len(shuffledDeck))
		}
	})

	t.Run("StringRepresentation gives correct player representation", func(t *testing.T) {
		hand := Hand{[]deck.Card{{10, 'H'}, {'Q', 'C'}}, player}
		expectedRepr := "10 of H, Q of C - value: 20"
		repr := hand.StringRepresentation()
		if repr != expectedRepr {
			t.Errorf("Expected %v got %v", expectedRepr, repr)
		}
	})

	t.Run("StringRepresentation gives correct dealer representation", func(t *testing.T) {
		hand := Hand{[]deck.Card{{10, 'H'}, {'Q', 'C'}}, dealer}
		expectedRepr := "10 of H, [hidden card]"
		repr := hand.StringRepresentation()
		if repr != expectedRepr {
			t.Errorf("Expected %v got %v", expectedRepr, repr)
		}
	})
}
