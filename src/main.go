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
	fmt.Printf("Your hand is %v\n", hand.StringRepresentation(playerHand))
}
