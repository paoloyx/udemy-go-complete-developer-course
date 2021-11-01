package main

import "fmt"

type deck []string

// print deck function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// returns a new complete deck of cards
func newDeck() deck {

	var cards deck
	var cardSuites = []string{"Cuori", "Picche", "Mattoni", "Cuori"}
	var cardValues = []string{"Uno", "Due", "Tre", "Quattro", "Cinque", "Sei", "Sette", "Otto", "Nove", "Dieci", "Fante", "Donna", "Re"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" di "+suite)
		}
	}

	return cards

}

// Deals a "hand" of cards from the deck
func deal(d deck, handSize int) (deck, deck) {

	hand := d[:handSize]
	remainingCards := d[handSize:]

	return hand, remainingCards

}
