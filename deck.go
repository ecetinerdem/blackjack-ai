package main

import (
	"fmt"
	"math/rand/v2"
)

type Deck []Card

func NewDeck() Deck {

	deck := Deck{}
	suits := []string{Hearts, Diamonds, Clubs, Spades}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	scores := []int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

	// Create a standart deck of 52 cards

	for _, suit := range suits {
		for i, value := range values {
			deck = append(deck, Card{
				Suit:  suit,
				Value: value,
				Score: scores[i],
			})

		}
	}
	return deck
}

func (d Deck) Shuffle() Deck {
	// Creadte a shuflled deck of cards

	shuffled := make(Deck, len(d))

	copy(shuffled, d)

	// Fisher-Yates shuffle algorithm
	for i := len(shuffled) - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled

}

func (d *Deck) Draw() Card {
	// If d empty..
	if len(*d) == 0 {
		fmt.Println("Deck is empty! Reshuffling...")
		*d = NewDeck().Shuffle()
	}

	card := (*d)[0]
	*d = (*d)[1:]
	return card
}
