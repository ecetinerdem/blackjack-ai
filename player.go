package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	Hit            = "h"
	Stand          = "s"
	Quit           = "q"
	MinDealerStand = 17
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

func (p *Player) handleHit(deck *Deck, cardCounter *CardCounter) bool {
	// Draw a card and add it to the player hand
	card := deck.Draw()
	p.AddCard(card, cardCounter)

	// Show the card that they got
	fmt.Printf("%s drew: %s\n", p.Name, card.String())
	p.DisplayHand(false)

	// Check if they went over 21
	if p.Score > 21 {
		fmt.Printf("%s busts with a score over 21!\n", p.Name)
		p.IsBust = true
		return true
	}

	// If it is the ai turn add a small delay to make it easier to follow

	if p.IsAI {
		time.Sleep(1 * time.Second)
	}

	return false
}

func (p *Player) PlayTurn(deck *Deck, cardCounter *CardCounter, dealerUpCard Card) {
	if p.IsAI {

	} else {
		// if it is a human let them choose what to do
		p.playHumanTurn(deck, cardCounter)
	}
}

func (p *Player) playHumanTurn(deck *Deck, cardCounter *CardCounter) {
	fmt.Printf("\n--- %s's Turn ---\n", p.Name)

	// Keep asking them what they want to do? (h)it, (s)tand,(q)uit
	fmt.Print("What would you like to do? (h)it, (s)tand,(q)uit: ")
	var choice string
	fmt.Scanln(&choice)
	choice = strings.ToLower(choice)

	switch choice {
	case Quit:
		fmt.Println("Thanks for playing!")
		return
	case Hit:
		if p.handleHit(deck, cardCounter) {
			return
		}
	case Stand:
		fmt.Printf("%s chose to stand.\n", p.Name)
		return
	default:
		fmt.Println("Invalid choice. Please try again")
	}
}
