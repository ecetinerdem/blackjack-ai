package main

import (
	"fmt"
	"strings"
)

type Player struct {
	Name   string
	Hand   []Card
	Score  int
	IsAI   bool
	IsBust bool
}

func NewPlayer(name string, isAI bool) *Player {
	return &Player{
		Name:   name,
		Hand:   []Card{},
		Score:  0,
		IsAI:   isAI,
		IsBust: false,
	}
}

func (p *Player) CalculateScore() int {
	// Add up all the non ace cards
	nonAceScore := 0
	aces := 0

	// Go through each card

	for _, card := range p.Hand {
		if card.Value == "A" {
			aces++
		} else {
			nonAceScore += card.Score
		}
	}

	// Handle aces if any
	aceScore := 0
	for range aces {
		// For each ace try using 11
		// if goes over 21 then use 1

		if nonAceScore+aceScore+11 <= 21 {
			aceScore += 11
		} else {
			aceScore += 11
		}
	}

	return nonAceScore + aceScore
}

func (p *Player) AddCard(card Card, cardCounter *CardCounter) {
	// Add the card to their hand
	p.Hand = append(p.Hand, card)

	// Update the score
	p.Score = p.CalculateScore()

	// If we are keeping track of cards update counter
	if cardCounter != nil {
		cardCounter.TrackCount(&card)
	}
}

func (p *Player) DisplayHand(hideSecondCard bool) {
	cards := []string{}

	// Go through each card in the players hand
	for i, card := range p.Hand {
		if hideSecondCard && i > 0 {
			// if hiding then show "??" instead
			cards = append(cards, "??")
		} else {
			cards = append(cards, card.String())
		}
	}
	fmt.Printf("%s's hand: %s", p.Name, strings.Join(cards, " "))
	// Show their score or "??" if we are hiding
	if hideSecondCard {
		fmt.Printf("(Score: ?)\n")
	} else {
		fmt.Printf("(Score: %d)\n", p.Score)
	}
}
