package hand

import "../deck"

type Hand []deck.Card

func DealHand(shuffledDeck deck.Deck) Hand {
	cards := shuffledDeck[:3]
	var hand Hand
	for _, card := range cards {
		hand = append(hand, card)
	}
	return hand
}
