package deck

import (
	"math/rand"
	"testing"
)

func TestDeck(t *testing.T) {
	rand.Seed(1)

	t.Run("orderedDeck creates first Card", func(t *testing.T) {
		got := orderedDeck()[0]
		want := Card{'A', 'C'}
		if got != want { t.Errorf("got %q want %q", got, want)}
	})

	t.Run("orderedDeck creates last Card", func(t *testing.T) {
		got := orderedDeck()[51]
		want := Card{'K', 'S'}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("shuffledDeck returns shuffled deck", func(t *testing.T) {
		expectedFirst := Card{10, 'H'}
		expectedLast := Card{8, 'S'}

		got := ShuffledDeck()

		actualFirst := got[0]
		actualLast := got[51]

		if actualFirst != expectedFirst {
			t.Errorf("Didn't get correct first card (expected %v actual %v)", expectedFirst, actualFirst)
		}
		if actualLast != expectedLast {
			t.Errorf("Didn't get correct last card (expected %v actual %v)", expectedLast, actualLast)
		}
	})
}
