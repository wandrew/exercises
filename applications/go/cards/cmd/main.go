package main

func main() {
	cards := newDeck()

	//deckLastIndex := len(cards) - 1

	cards.print()

	hand, cards := deal(cards, 3)

	hand.print()
	cards.print()
}

type Card string

func newCard() string {
	return "Five of Diamonds"
}
