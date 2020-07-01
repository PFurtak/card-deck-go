package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck len of 51, but received %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but recieved %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected the last card to be King onf Clubs, but received %v", d[len(d)-1])
	}
}
