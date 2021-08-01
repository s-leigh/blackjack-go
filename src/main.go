package main

import (
	"./deck"
	"./hand"
	"fmt"
	"io"
	"os"
	"strings"
)

const dealerSticksOnValue = 17

func main() {
	fmt.Println("Welcome to Blackjack!")

	shuffledDeck := deck.ShuffledDeck()
	playerHand := hand.DealHand(&shuffledDeck, false)
	dealerHand := hand.DealHand(&shuffledDeck, true)

	fmt.Printf("Your hand: %v\n", playerHand.PlayerStringRepresentation())
	fmt.Printf("Dealer's hand: %v\n", dealerHand.DealerStringRepresentation(true))

	playerGameLoop(&shuffledDeck, &playerHand)
	dealerGameLoop(&shuffledDeck, &dealerHand)
	checkWinCondition(playerHand, dealerHand, os.Stdout)
}

func playerGameLoop(shuffledDeck *deck.Deck, playerHand *hand.Hand) {
	var input string
	for strings.ToLower(input) != "s" {
		fmt.Println("[H]it or [S]tick?")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error accepting input, terminating...")
			return
		}

		if strings.ToLower(input) == "h" {
			playerHand.DealCard(shuffledDeck)
			fmt.Printf("Your hand: %v\n", playerHand.PlayerStringRepresentation())
			if isGameOver(playerHand) {
				return
			}
		}
	}
}

func dealerGameLoop(shuffledDeck *deck.Deck, dealerHand *hand.Hand) {
	fmt.Printf("Dealer's hand: %v\n", dealerHand.DealerStringRepresentation(false))
	for dealerHand.Value() < dealerSticksOnValue {
		dealerHand.DealCard(shuffledDeck)
		fmt.Printf("Dealer's hand: %v\n", dealerHand.DealerStringRepresentation(false))
	}
}

func isGameOver(hand *hand.Hand) bool {
	return hand.Value() >= 21
}

func checkWinCondition(playerHand, dealerHand hand.Hand, output io.Writer) {
	playerScore := playerHand.Value()
	dealerScore := dealerHand.Value()
	if playerScore == 21 {
		fmt.Fprintln(output, "Blackjack! You win!")
	} else if playerScore > 21 {
		fmt.Fprintln(output, "Bust!")
	} else {
		if playerScore > dealerScore {
			fmt.Fprintln(output, "You win!")
		} else {
			fmt.Fprintln(output, "You lose!")
		}
	}
	fmt.Fprintf(output, "Final scores:\nDealer: %d, You: %d\n", dealerScore, playerScore)
}
