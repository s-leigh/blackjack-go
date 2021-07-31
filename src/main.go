package main

import (
	"./deck"
	"./hand"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Blackjack!")

	shuffledDeck := deck.ShuffledDeck()
	playerHand := hand.DealHand(&shuffledDeck)
	dealerHand := hand.DealHand(&shuffledDeck)

	fmt.Printf("Your hand: %v\n", hand.StringRepresentation(playerHand, true))
	fmt.Printf("Dealer's hand: %v\n", hand.StringRepresentation(dealerHand, false))
}
