package main

import (
	"./deck"
	"./hand"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Blackjack!")

	shuffledDeck := deck.ShuffledDeck()
	playerHand := hand.DealHand(&shuffledDeck, false)
	dealerHand := hand.DealHand(&shuffledDeck, true)

	fmt.Printf("Your hand: %v\n", playerHand.StringRepresentation())
	fmt.Printf("Dealer's hand: %v\n", dealerHand.StringRepresentation())
}
