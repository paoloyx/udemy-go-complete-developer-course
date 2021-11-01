package main

func main() {
	// declare a slice
	cards := newDeck()
	cards.shuffle()

	// save deck to file, reads and print it
	cards.saveToFile("my_cards")
	cards = newDeckFromFile("my_cards")
	cards.print()
}
