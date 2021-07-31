package hand

import (
	"../deck"
	"fmt"
)

const handSize = 2

type owner int

const (
	dealer owner = iota
	player
)

type Hand struct {
	cards []deck.Card
	owner owner
}

func (h Hand) Value() int {
	value := 0
	for _, card := range h.cards {
		if card.Rank <= 10 {
			value += int(card.Rank)
		} else if card.Rank != 'A' {
			value += 10
		}
	}
	// calculate aces after all other cards added up
	numberOfAces := func() int {
		count := 0
		for _, card := range h.cards {
			if card.Rank == 'A' {
				count++
			}
		}
		return count
	}()

	if numberOfAces == 0 {
		return value
	}

	if value+numberOfAces < 11 {
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
	hand.cards = cards
	if isDealer {
		hand.owner = dealer
	} else {
		hand.owner = player
	}
	return hand
}

func (h Hand) StringRepresentation() string {
	var stringRep string
	var numCardsToShow int
	if h.owner == player {
		numCardsToShow = len(h.cards)
	} else {
		numCardsToShow = 1
	}
	for i := 0; i < numCardsToShow; i++ {
		card := h.cards[i]
		var value string
		if card.Rank > 10 {
			value = string(card.Rank)
		} else {
			value = fmt.Sprintf("%d", card.Rank)
		}
		stringRep += fmt.Sprintf("%v of %v", value, string(card.Suit))
		if i != len(h.cards)-1 {
			stringRep += ", "
		}
		if h.owner == dealer {
			stringRep += "[hidden card]"
		}
	}
	if h.owner != dealer {
		stringRep += fmt.Sprintf(" - value: %d", h.Value())
	}
	return stringRep
}
