package hand

import (
	"../deck"
)

const handSize = 2

type Hand []deck.Card

func DealHand(shuffledDeck *deck.Deck) Hand {
	cards := (*shuffledDeck)[:handSize-1]
	*shuffledDeck = (*shuffledDeck)[handSize:]
	var hand Hand
	for _, card := range cards {
		hand = append(hand, card)
	}
	return hand
}
