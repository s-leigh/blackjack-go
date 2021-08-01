package game

import (
	"../deck"
	"../hand"
	"bytes"
	"math/rand"
	"testing"
)

func TestGameLogic(t *testing.T) {
	rand.Seed(1)

	t.Run("isGameOver returns true when >=21 or false otherwise", func(t *testing.T) {
		gameOverHand := hand.Hand{[]deck.Card{deck.Card{'Q', 'S'}, deck.Card{'Q', 'D'}, deck.Card{'Q', 'C'}}}
		if isGameOver(&gameOverHand) != true {
			t.Errorf("Expected isGameOver to return true")
		}
		notGameOverHand := hand.Hand{[]deck.Card{{'Q', 'S'}, {2, 'S'}}}
		if isGameOver(&notGameOverHand) != false {
			t.Errorf("Expected isGameOver to return false")
		}
	})

	t.Run("checkWinCondition identifies player's 21 score", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		playerHand := hand.Hand{Cards: []deck.Card{{'A', 'S'}, {'Q', 'D'}}}
		dealerHand := hand.Hand{Cards: []deck.Card{{'A', 'D'}, {'Q', 'S'}}}
		CheckWinCondition(playerHand, dealerHand, buffer)
		result, _ := buffer.ReadString([]byte("\n")[0])
		if result != "Blackjack! You win!\n" {
			t.Errorf("Unexpected string: %v", result)
		}
	})
	t.Run("checkWinCondition identifies player's bust score", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		playerHand := hand.Hand{Cards: []deck.Card{{'Q', 'C'}, {'Q', 'D'}, {'Q', 'S'}}}
		dealerHand := hand.Hand{Cards: []deck.Card{{'A', 'D'}, {'Q', 'S'}}}
		CheckWinCondition(playerHand, dealerHand, buffer)
		result, _ := buffer.ReadString([]byte("\n")[0])
		if result != "Bust!\n" {
			t.Errorf("Unexpected string: %v", result)
		}
	})
	t.Run("checkWinCondition identifies player's winning score", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		playerHand := hand.Hand{Cards: []deck.Card{{'Q', 'C'}, {'Q', 'D'}}}
		dealerHand := hand.Hand{Cards: []deck.Card{{3, 'D'}, {'Q', 'S'}}}
		CheckWinCondition(playerHand, dealerHand, buffer)
		result, _ := buffer.ReadString([]byte("\n")[0])
		if result != "You win!\n" {
			t.Errorf("Unexpected string: %v", result)
		}
	})
	t.Run("checkWinCondition identifies dealer's winning score on tie", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		playerHand := hand.Hand{Cards: []deck.Card{{'Q', 'C'}, {'Q', 'D'}}}
		dealerHand := hand.Hand{Cards: []deck.Card{{'Q', 'D'}, {'Q', 'S'}}}
		CheckWinCondition(playerHand, dealerHand, buffer)
		result, _ := buffer.ReadString([]byte("\n")[0])
		if result != "You lose!\n" {
			t.Errorf("Unexpected string: %v", result)
		}
	})
	t.Run("dealerGameLoop deals cards until dealer's stick limit", func(t *testing.T) {
		shuffledDeck := deck.ShuffledDeck()
		dealerHand := hand.Hand{Cards: []deck.Card{}}
		// two non-randomised cards give value of 17, the stick limit
		DealerGameLoop(&shuffledDeck, &dealerHand)
		dealerHandSize := len(dealerHand.Cards)
		if dealerHandSize != 2 {
			t.Errorf("Expected hand to contain 2 cards, got %d", dealerHandSize)
		}
		deckSize := len(shuffledDeck)
		if deckSize != 50 {
			t.Errorf("Expected deck to contain 50 cards, got %d", deckSize)
		}
	})
}
