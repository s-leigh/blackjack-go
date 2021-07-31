package main

import (
	"./deck"
	"./hand"
	"testing"
)

func TestMain(t *testing.T) {
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

	// TODO other tests
}
