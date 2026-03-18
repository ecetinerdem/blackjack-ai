package main

import (
	"fmt"
	"strings"
)

func main() {
	// Clear the Screen
	clearScreen()
	// Welcome message
	fmt.Println("=== Go BlackJack ===")
	fmt.Println("Welcome to Go BlackJack! You are playing against an AI with card counting abilities")
	// Get a deck of cards
	deck := NewDeck().Shuffle()

	for _, c := range deck {
		fmt.Printf("%s%s\n", c.Value, c.Suit)
	}
	// Create a card counter for the AI to use
	cardCounter := NewCardCounter()
	fmt.Println(cardCounter.DecksRemaining)

	for {
		// Check to see if we need to shuffle the deck
		if len(deck) < 10 {
			fmt.Println("\n=== Decik is running low. Reshuffling... ===")
			deck = NewDeck().Shuffle()
			cardCounter.Reset()
			fmt.Println("Deck reshuffled and card counter reset.")
		}
		// Play a round
		PlayRound(&deck, cardCounter)

		// Ask if the player wants to play another round
		fmt.Printf("\n Play another round? (y/n): ")
		var choice string
		fmt.Scanln(&choice)
		choice = strings.ToLower(choice)

		// If not quit the game

		if choice == "n" {
			fmt.Println("Thanks for playing!")
			break
		}

	}
}
