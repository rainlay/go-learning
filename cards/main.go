package main

func main() {
	cards := newDeckFromFile("cardsFile")
	//cards := newDeckFromFile("nofile")
	cards.print()
}

