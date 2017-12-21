package main

func main() {
	cards := newDeck()

	hand, remainCards := deal(cards, 5)

	hand.print()
	remainCards.print()
}
