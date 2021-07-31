package hand

import (
	"../deck"
	"fmt"
)

const handSize = 2

type Hand []deck.Card

func (h Hand) Value() int {
	value := 0
	for _, card := range h {
		if card.Rank <= 10 {
			value += int(card.Rank)
		} else if card.Rank != 'A' {
			value += 10
		}
	}
	// calculate aces after all other cards added up
	numberOfAces := func() int {
		count := 0
		for _, card := range h {
			if card.Rank == 'A' {
				count++
			}
		}
		return count
	}()

	if numberOfAces == 0 {
		return value
	}

	if value + numberOfAces < 11 {
		value += 10 + numberOfAces
	} else {
		value += numberOfAces
	}

	return value
}

func DealHand(shuffledDeck *deck.Deck) Hand {
	cards := (*shuffledDeck)[:handSize]
	*shuffledDeck = (*shuffledDeck)[handSize:]
	var hand Hand
	for _, card := range cards {
		hand = append(hand, card)
	}
	return hand
}

func StringRepresentation(hand Hand, playersHand bool) string {
	var stringRep string
	var numCardsToShow int
	if playersHand {
		numCardsToShow = len(hand)
	} else {
		numCardsToShow = 1
	}
	for i := 0; i < numCardsToShow; i++ {
		card := hand[i]
		var value string
		if card.Rank > 10 {
			value = string(card.Rank)
		} else {
			value = fmt.Sprintf("%d", card.Rank)
		}
		stringRep += fmt.Sprintf("%v of %v", value, string(card.Suit))
		if i != len(hand)-1 {
			stringRep += ", "
		}
		if !playersHand {
			stringRep += "[hidden card]"
		}
	}
	return stringRep
}
