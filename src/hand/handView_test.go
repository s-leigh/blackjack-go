package hand

import (
	"../deck"
	"testing"
)

func TestHandView(t *testing.T) {
	t.Run("StringRepresentation gives correct player representation", func(t *testing.T) {
		hand := Hand{[]deck.Card{{10, 'H'}, {'Q', 'C'}}}
		expectedRepr := "10 of H, Q of C - value: 20"
		repr := hand.PlayerStringRepresentation()
		if repr != expectedRepr {
			t.Errorf("Expected %v got %v", expectedRepr, repr)
		}
	})

	t.Run("StringRepresentation gives correct dealer representation with hidden card", func(t *testing.T) {
		hand := Hand{[]deck.Card{{10, 'H'}, {'Q', 'C'}}}
		expectedRepr := "10 of H, [hidden card]"
		repr := hand.DealerStringRepresentation(true)
		if repr != expectedRepr {
			t.Errorf("Expected %v got %v", expectedRepr, repr)
		}
	})

	t.Run("StringRepresentation gives correct dealer representation without hidden card", func(t *testing.T) {
		hand := Hand{[]deck.Card{{10, 'H'}, {'Q', 'C'}}}
		expectedRepr := "10 of H, Q of C - value: 20"
		repr := hand.DealerStringRepresentation(false)
		if repr != expectedRepr {
			t.Errorf("Expected %v got %v", expectedRepr, repr)
		}
	})
}
