package main

import "fmt"

func main() {
	cards, err := newDeckFromFile("my_cards")
	if err != nil {
		fmt.Println("An error was encountered. Error: ", err)
		fmt.Println("Creating a new Deck")
		cards = newDeck()
	} else {
		fmt.Println("A saved deck was found")
	}
	cards.print()
	fmt.Println("Shuffling")
	cards = shuffle(cards)
	cards = shuffle(cards)
	cards = shuffle(cards)
	cards.print()
}
