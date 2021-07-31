package hand

import (
	"../deck"
	"fmt"
)

const handSize = 2

type Hand []deck.Card

func DealHand(shuffledDeck *deck.Deck) Hand {
	cards := (*shuffledDeck)[:handSize]
	*shuffledDeck = (*shuffledDeck)[handSize:]
	var hand Hand
	for _, card := range cards {
		hand = append(hand, card)
	}
	return hand
}

func StringRepresentation(hand Hand) string {
	var stringRep string
	for i, card := range hand {
		var value string
		if card.Value > 10 {
			value = string(card.Value)
		} else {
			value = fmt.Sprintf("%d", card.Value)
		}
		stringRep += fmt.Sprintf("%v of %v", value, string(card.Suit))
		if i != len(hand) - 1 {
			stringRep += ", "
		}
	}
	return stringRep
}
