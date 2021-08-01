package main

import (
	"./deck"
	"./game"
	"./hand"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Blackjack!")

	shuffledDeck := deck.ShuffledDeck()
	playerHand := hand.DealHand(&shuffledDeck, false)
	dealerHand := hand.DealHand(&shuffledDeck, true)

	fmt.Printf("Your hand: %v\n", playerHand.PlayerStringRepresentation())
	fmt.Printf("Dealer's hand: %v\n", dealerHand.DealerStringRepresentation(true))

	game.PlayerGameLoop(&shuffledDeck, &playerHand)
	game.DealerGameLoop(&shuffledDeck, &dealerHand)
	game.CheckWinCondition(playerHand, dealerHand, os.Stdout)
}
