package main

import (
	"fmt"
	"time"
)

func AdvancedAIDecision(player Player, cardCounter *CardCounter, dealerUpCard Card) string {
	score := player.Score

	if score >= 19 {
		return Stand
	}

	// Calculate bust probability if we hit
	bustProbability := cardCounter.ChanceofBusting(score)

	// Calculate Dealer bust probability
	dealerBustProbability := cardCounter.DealerChanceOfBusting(dealerUpCard)

	// Display AI thinking
	fmt.Printf("AI thinking: Score %d, True Count %.1f, Bust probability %.1f%%, Dealer Bust probability %.1f%%\n", score, cardCounter.TrueCount, bustProbability*100, dealerBustProbability*100)
	time.Sleep(500 * time.Millisecond)

	// Decision logic incorporating card counting
	if score >= 17 {
		// With 17 and 18 consider the true count
		if cardCounter.TrueCount > 0 {
			// More high cards
			return Stand
		} else if dealerUpCard.Score >= 7 && cardCounter.TrueCount < -2 {
			// Against a strong dealer with negative count so hit on 17
			if score == 17 {
				return Hit
			}
			return Stand
		}
		return Stand
	}
	// Soft hands (have an ace counted as 11)
	hasAce := false
	for _, card := range player.Hand {
		if card.Value == "A" && score <= 21 {
			hasAce = true
			break
		}
	}

	if hasAce {
		if score >= 18 {
			if (dealerUpCard.Score >= 9 || dealerUpCard.Value == "A") && cardCounter.TrueCount < -1 {
				return Hit
			}
			return Stand
		}

		return Hit
	}

	if score <= 16 {
		// if low risking of busting and dealer likely to bust, stand
		if bustProbability < 0.3 && dealerBustProbability > 0.4 && score >= 13 {
			return Stand
		}

		if score >= 12 && dealerUpCard.Score >= 2 && dealerUpCard.Score <= 6 && cardCounter.TrueCount > -3 {
			return Stand
		}
		return Hit
	}
	return Stand
}
