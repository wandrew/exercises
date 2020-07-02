package main

import "fmt"

// create a custom type for a deck of cards
type deck []string
type card struct {
	suit string
	rank string
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardRank := []string{"Ace", "King", "Queen", "Jack"}
	for _, suit := range cardSuits {
		for _, value := range cardRank {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	deck := d[handSize:]

	return hand, deck
}
