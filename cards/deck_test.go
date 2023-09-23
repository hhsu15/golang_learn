package main

import "testing"

// change from
// t.Errorf("Expected deck length of 16, but got", len(d))
// to
// t.Errorf("Expected dexk length og 16, but got %v", len(d))
func TestNewDeck (t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First card should be Ace of Spades but got %v", d[0])
	}
	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Last card should ne Four of Clubs but got %v", d[len(d) - 1])
	}
}