package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// returns a toString implementation of the deck
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Writes a deck into a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Loads a deck from a file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Shuffles a deck
func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
