package main

import (
	"./deck"
	"./hand"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to Blackjack!")

	shuffledDeck := deck.ShuffledDeck()
	playerHand := hand.DealHand(&shuffledDeck, false)
	dealerHand := hand.DealHand(&shuffledDeck, true)

	fmt.Printf("Your hand: %v\n", playerHand.StringRepresentation())
	fmt.Printf("Dealer's hand: %v\n", dealerHand.StringRepresentation())

	var input string
	fmt.Println("[H]it or [S]tick?")

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error accepting input, terminating...")
		return
	}

	if strings.ToLower(input) == "h" {
		playerHand.DealCard(&shuffledDeck)
		fmt.Printf("Your hand: %v\n", playerHand.StringRepresentation())
	}
}
