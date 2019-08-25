package main

func main() {
	//cards := newDeckFromFile("cardsFile")
	//cards := newDeckFromFile("nofile")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

