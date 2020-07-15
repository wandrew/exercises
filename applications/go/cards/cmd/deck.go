package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// create a custom type for a deck of cards
type deck []card
type card struct {
	suit string
	rank string
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardRank := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardRank {
			c := card{suit: suit, rank: value}
			cards = append(cards, c)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, s int) (deck, deck) {
	// [:s] reaches from index 0 up to, but not including s
	// [s:] starts from and includes s and includes everything through the last index
	return d[:s], d[s:]
}

func shuffle(d deck) deck {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// This was the course approach
	for i := range d {
		n := r.Intn(len(d) - 1)
		d[i], d[n] = d[n], d[i]
	}

	return d
}

// this was my approach
func shuffleB(d deck) deck {
	//src := rand.NewSource(time.Now().UnixNano())
	//r := rand.New(src)

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})

	return d
}

func (d deck) toString() string {
	// Join requires a slice of strings, our "deck" type is a slice of strings.  A type conversion is unnecessary
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	if err != nil {
		return err
	}
	return nil
}

func newDeckFromFile(filename string) (deck, error) {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return deck(strings.Split(string(in), ",")), nil
}
