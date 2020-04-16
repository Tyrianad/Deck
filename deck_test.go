package main

import "testing"

func TestAddCard(t *testing.T) {
	d := deck{}
	d.AddCard(card{"A", "♤"})

	if len(d.cards) != 1 {
		t.Error("Expected 1 card, got: ", len(d.cards))
	}
}
