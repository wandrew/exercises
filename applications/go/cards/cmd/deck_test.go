package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	got := string(len(d))
	want := string(52)
	if got != want {
		t.Errorf("Expected deck lenth of %v but got %v", want, got)
	}

	got = d[0]
	want = "Ace of Spades"
	if got != want {
		t.Errorf("Expected first card to be %v but got %v", want, got)
	}

	got = d[len(d)-1]
	want = "King of Diamonds"
	if got != want {
		t.Errorf("Expected last card to be %v but got %v", want, got)
	}

}
