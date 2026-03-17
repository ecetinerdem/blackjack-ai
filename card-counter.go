package main

const deckSize = 52.0

type CardCounter struct {
	SeenCard       map[string]int // Kep track of seen cards by value
	RunningCount   int            // Runninng count for hi-lo strategy
	TrueCount      float64        // True count(running count / cards remaining)
	DecksRemaining float64        // Estimate of remaining cards
}

func NewCardCounter() *CardCounter {
	counter := &CardCounter{
		SeenCard:       make(map[string]int),
		RunningCount:   0,
		TrueCount:      0,
		DecksRemaining: 1.0,
	}

	// Initialize counts for each card value 0

	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, value := range values {
		counter.SeenCard[value] = 0
	}
	return counter
}

func (cc *CardCounter) Reset() {
	cc.RunningCount = 0
	cc.TrueCount = 0
	cc.DecksRemaining = 1.0
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, value := range values {
		cc.SeenCard[value] = 0
	}
}

// TrackCard update card counter's state based on newly seen card
func (cc *CardCounter) TrackCount(card *Card) {
	// Update seen card state
	cc.SeenCard[card.Value]++

	// Update the running card using the hi-lo
	if card.Value == "2" || card.Value == "3" || card.Value == "4" || card.Value == "5" || card.Value == "6" {
		cc.RunningCount++
	} else if card.Value == "10" || card.Value == "J" || card.Value == "Q" || card.Value == "K" {
		cc.RunningCount--
	}

	// Update decks remaining estimation (52 cards in deck)
	totalSeen := 0

	for _, count := range cc.SeenCard {
		totalSeen += count
	}

	cc.DecksRemaining = (52.0 - float64(totalSeen)) / deckSize
	if cc.DecksRemaining < 0.1 {
		cc.DecksRemaining = 0.1 // avoid division by too small of a number
	}

	// Calculate the true count
	cc.TrueCount = float64(cc.RunningCount) / cc.DecksRemaining
}

func (cc *CardCounter) ChanceofBusting(playerScore int) float64 {
	// If player has 21 or more they will bust
	if playerScore >= 21 {
		return 1.0
	}

	// Calculate how many points untill player busts
	pointsUntilBust := 21 - playerScore

	// Count unseen cards that would cause a bust
	bustCards := 0
	totalUnSeenCards := 0

	// For each card value check if it would cause a bust
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	scores := []int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

	for i, value := range values {
		totalDeck := 4
		seen := cc.SeenCard[value]
		unseen := totalDeck - seen
		if unseen < 0 {
			unseen = 0
		}
		totalUnSeenCards += unseen

		if scores[i] > pointsUntilBust {
			bustCards += unseen
		}
	}
	if totalUnSeenCards == 0 {
		return 0.5 // default %50 if we somehow have unseen cards
	}
	return float64(bustCards) / float64(totalUnSeenCards)

}

func (cc *CardCounter) DealerChanceOfBusting(dealerUpCard Card) float64 {
	// Base probabilities of dealer busting based on uon up card
	bustProbabilities := map[string]float64{
		"A":  0.17,
		"2":  0.35,
		"3":  0.37,
		"4":  0.40,
		"5":  0.42,
		"6":  0.42,
		"7":  0.26,
		"8":  0.24,
		"9":  0.23,
		"10": 0.21,
		"J":  0.21,
		"Q":  0.21,
		"K":  0.21,
	}

	// Adjust base probability based on our card counting
	baseProbability := bustProbabilities[dealerUpCard.Value]

	// if true countis positive (more low cards have been seen)
	adjustment := cc.TrueCount * 0.02

	return baseProbability + adjustment
}
