package main

func main() {
	// declare a slice
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

}
