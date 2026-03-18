package main

import "fmt"

func PlayRound(deck *Deck, cardCounter *CardCounter) {
	fmt.Println("\n=== New Round of BlackJack ===")
	fmt.Printf("Cards remaining in deck: %d\n", len(*deck))

	// Initialize players
	dealer := NewPlayer("Dealer", false)
	human := NewPlayer("Human", false)
	ai := NewPlayer("AI", true)

	fmt.Println("Players: ", dealer.Name, human.Name, ai.Name)

	// Initial deal: Two cards per player
	for range 2 {
		human.AddCard(deck.Draw(), cardCounter)
		ai.AddCard(deck.Draw(), cardCounter)
		dealer.AddCard(deck.Draw(), cardCounter)
	}

	// Show initial hands
	fmt.Println("\nInitial Deal:")
	dealer.DisplayHand(true)
	human.DisplayHand(false)
	ai.DisplayHand(false)

	// Play each player's turn
	human.PlayTurn(deck, cardCounter, dealer.Hand[0])

	// Show results

	// Display results

	// Display card counting statistics
}
