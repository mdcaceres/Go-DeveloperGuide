package main

import "fmt"

func main() {
	//var card string = "Ace of spades"
	//card := "Ace of Spades"
	//card := newCard()
	//fmt.Println(card)
	//cards := deck{newCard(), newCard()}
	//cards = append(cards, "Ace of Spades")
	//fmt.Println(cards)

	//fmt.Println("que carajos?")

	//cards := newDeck()

	//fmt.Println(cards.toString())

	//cards.print()

	//hand, remainingCards := deal(cards, 5)

	//hand.print()

	//remainingCards.print()

	//greeting := "Hello there!"

	//type conversion
	//fmt.Println([]byte(greeting))

	//cards.saveToFile("my_cards2")

	cards := newDeckFromFile("my_cards")
	fmt.Println(cards)
}
