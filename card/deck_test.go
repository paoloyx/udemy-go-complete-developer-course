package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected new deck lenght of %d but got %d", 52, len(d))
	}

}
