package hand

import (
	"../deck"
	"math/rand"
	"testing"
)

func TestHand(t *testing.T) {
	rand.Seed(1)

	t.Run("Hand value without aces calculated correctly", func(t *testing.T) {
		hand1 := Hand{[]deck.Card{{2, 'H'}, {2, 'D'}}}
		if hand1.Value() != 4 {
			t.Errorf("Expected hand value to be 4, got %d", hand1.Value())
		}

		handWithFaceCard := Hand{[]deck.Card{{'Q', 'H'}, {2, 'D'}}}
		if handWithFaceCard.Value() != 12 {
			t.Errorf("Expected hand value to be 12, got %d", handWithFaceCard.Value())
		}
	})

	t.Run("Hand value with aces calculated correctly", func(t *testing.T) {
		handWithAceHigh := Hand{[]deck.Card{{'A', 'S'}, {'Q', 'D'}}}
		if handWithAceHigh.Value() != 21 {
			t.Errorf("Expected hand value to be 21, got %d", handWithAceHigh.Value())
		}

		handWithAceLow := Hand{[]deck.Card{{'A', 'H'}, {'Q', 'D'}, {4, 'D'}}}
		if handWithAceLow.Value() != 15 {
			t.Errorf("Expected hand value to be 15, got %d", handWithAceLow.Value())
		}

		handWithTwoAces := Hand{[]deck.Card{{'A', 'H'}, {'A', 'C'}}}
		if handWithTwoAces.Value() != 12 {
			t.Errorf("Expected hand value to be 12, got %d", handWithTwoAces.Value())
		}

		handWithThreeAces := Hand{[]deck.Card{{'A', 'H'}, {'A', 'C'}, {'A', 'D'}}}
		if handWithThreeAces.Value() != 13 {
			t.Errorf("Expected hand value to be 13, got %d", handWithThreeAces.Value())
		}

	})

	t.Run("DealHand returns hand", func(t *testing.T) {
		shuffledDeck := deck.ShuffledDeck()
		expectedHand := Hand{[]deck.Card{{Rank: 10, Suit: 'H'}, {Rank: 7, Suit: 'H'}, {Rank: 'J', Suit: 'S'}, {Rank: 'Q', Suit: 'S'}}}
		hand := DealHand(&shuffledDeck, false)

		if len(hand.Cards) != 2 {
			t.Errorf("hand length is %v", len(hand.Cards))
		}
		for i, card := range hand.Cards {
			if card != expectedHand.Cards[i] {
				t.Errorf("got %v want %v", card, expectedHand.Cards[i])
			}
		}
	})

	t.Run("DealHand removes cards from deck", func(t *testing.T) {
		shuffledDeck := deck.ShuffledDeck()
		DealHand(&shuffledDeck, false)
		DealHand(&shuffledDeck, true)
		expectedDeckSize := len(shuffledDeck)
		if expectedDeckSize != 48 {
			t.Errorf("Expected deck size of 48, got %v", expectedDeckSize)
		}
	})

	t.Run("DealCard adds card to hand", func(t *testing.T) {
		shuffledDeck := deck.ShuffledDeck()
		hand := Hand{[]deck.Card{{10, 'H'}, {'Q', 'C'}}}
		hand.DealCard(&shuffledDeck)
		if len(hand.Cards) != 3 {
			t.Errorf("Expected a hand of length 3, got %d", len(hand.Cards))
		}
		if len(shuffledDeck) != 51 {
			t.Errorf("Expected deck size of 51, got %d", len(shuffledDeck))
		}
	})
}
