package hand

import (
	"../deck"
	"fmt"
)

func (h Hand) DealerStringRepresentation(onlyShowFirstCard bool) string {
	if onlyShowFirstCard {
		return cardStringRepresentation([]deck.Card{h.Cards[0]}) + ", [hidden card]"
	} else {
		return cardStringRepresentation(h.Cards) + fmt.Sprintf(" - value: %d", h.Value())
	}
}

func (h Hand) PlayerStringRepresentation() string {
	return cardStringRepresentation(h.Cards) + fmt.Sprintf(" - value: %d", h.Value())
}

func cardStringRepresentation(cards []deck.Card) string {
	var stringRep string
	for i, card := range cards {
		var value string
		if card.Rank > 10 {
			value = string(card.Rank)
		} else {
			value = fmt.Sprintf("%d", card.Rank)
		}
		stringRep += fmt.Sprintf("%v of %v", value, string(card.Suit))
		if i != len(cards)-1 {
			stringRep += ", "
		}
	}
	return stringRep
}
