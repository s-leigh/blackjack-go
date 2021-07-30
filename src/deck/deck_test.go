package deck

import "testing"

func TestDeck(t *testing.T) {
	t.Run("orderedDeck creates first card", func(t *testing.T) {
		got := orderedDeck()[0]
		want := card{'A', 'C'}
		if got != want { t.Errorf("got %q want %q", got, want)}
	})

	t.Run("orderedDeck creates last card", func(t *testing.T) {
		got := orderedDeck()[51]
		want := card{'K', 'S'}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("shuffledDeck returns correct length of deck", func(t *testing.T) {
		got := ShuffledDeck()
		if len(got) != 52 {
			t.Errorf("Didn't get correct deck length")
		}
	})
}
