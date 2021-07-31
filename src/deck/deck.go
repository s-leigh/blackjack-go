package deck

import (
	"math/rand"
	"time"
)

type card struct {
	value rune
	suit rune
}

type Deck []card

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
	for _, value := range values {
		for _, suit := range suits {
			card := card{value, suit}
			orderedDeck = append(orderedDeck, card)
		}
	}
	return orderedDeck
}
