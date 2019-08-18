package main

func main()  {
	// var card string = "Ace of Spades"
	//cards := []string{newCard(),newCard()}
	cards := deck{"Ace of Diamonds",newCard()}
	cards = append(cards,"Six of Spades")

	cards.print()
}

func newCard() string{
	return "Five of Diamonds"
}