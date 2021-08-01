package hand

import (
	"../deck"
)

const handSize = 2

type Hand struct {
	Cards []deck.Card
}

func (h Hand) Value() int {
	value := 0
	for _, card := range h.Cards {
		if card.Rank <= 10 {
			value += int(card.Rank)
		} else if card.Rank != 'A' {
			value += 10
		}
	}
	// calculate aces after all other cards added up
	numberOfAces := func() int {
		count := 0
		for _, card := range h.Cards {
			if card.Rank == 'A' {
				count++
			}
		}
		return count
	}()

	if numberOfAces == 0 {
		return value
	}

	if value+numberOfAces <= 11 {
		value += 10 + numberOfAces
	} else {
		value += numberOfAces
	}

	return value
}

func DealHand(shuffledDeck *deck.Deck, isDealer bool) Hand {
	cards := (*shuffledDeck)[:handSize]
	*shuffledDeck = (*shuffledDeck)[handSize:]
	var hand Hand
	hand.Cards = cards
	return hand
}

func (h *Hand) DealCard(shuffledDeck *deck.Deck) {
	newCard := (*shuffledDeck)[:1]
	*shuffledDeck = (*shuffledDeck)[1:]
	h.Cards = append(h.Cards, newCard[0])
}
