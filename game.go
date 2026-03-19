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

	if !human.IsBust {
		// Let ai player play
		ai.PlayTurn(deck, cardCounter, dealer.Hand[0])
		dealer.playDealerTurn(deck, cardCounter)
	}

	// Show results
	fmt.Println("\n=== Results ===")
	fmt.Printf("Dealer: %d\n", dealer.Score)
	fmt.Printf("Human: %d\n", human.Score)
	fmt.Printf("AI: %d\n", ai.Score)

	// Display results
	fmt.Println(human.DetermineResult(*dealer))
	fmt.Println(ai.DetermineResult(*dealer))

	// Display card counting statistics

	displayCardCountingStats(deck, cardCounter)
}

func displayCardCountingStats(deck *Deck, cardCounter *CardCounter) {
	fmt.Println("\n=== Card Counting Statistics ===")
	fmt.Printf("Final Running Count: %d\n", cardCounter.RunningCount)
	fmt.Printf("Final True Count: %.f\n", cardCounter.TrueCount)
	fmt.Printf("Cards R emaining in Deck: %d\n", len(*deck))

	fmt.Printf("\n Card Distribution Seen: ")
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, value := range values {
		fmt.Printf("%s: %d   ", value, cardCounter.SeenCard[value])
		if value == "6" {
			fmt.Println() // Put in a line break for readebility
		}
	}
	fmt.Println()
}
