package deck

import (
	"math/rand"
	"time"
)

type card struct {
	value rune
	suit rune
}

type Deck [52]card

func ShuffledDeck() Deck {
	rand.Seed(time.Now().UnixNano())
	deck := orderedDeck()
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}

func orderedDeck() Deck {
	var values = [13]rune{'A', 2, 3, 4, 5, 6, 7, 8, 9, 10, 'J', 'Q', 'K'}
	var suits = [4]rune{'C', 'D', 'H', 'S'}
	var orderedDeck Deck
	for valueIndex, value := range values {
		for suitIndex, suit := range suits {
			card := card{value, suit}
			deckIndex := valueIndex + (suitIndex * len(values))
			orderedDeck[deckIndex] = card
		}
	}
	return orderedDeck
}
