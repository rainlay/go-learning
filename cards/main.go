package main

import "fmt"

func main()  {
	// var card string = "Ace of Spades"
	//cards := []string{newCard(),newCard()}
	cards := deck{"Ace of Diamonds",newCard()}
	cards = append(cards,"Six of Spades")

	// like php foreach
	// i => index
	// card => current iterating cards value
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string{
	return "Five of Diamonds"
}