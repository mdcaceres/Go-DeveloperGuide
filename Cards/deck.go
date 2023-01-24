package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create new type of deck
type deck []string

// append card suits and values to cards slice and return it
func newDeck() deck {
	cards := deck{}
	//Slice of string
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// return multiple values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	//if nothing went wrong err will have nil (NIL) == VALUE WICH MEANS NO VALUE!

	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {
	//seed of generator
	source := rand.NewSource(time.Now().UnixNano())
	// random generator
	r := rand.New(source)

	//shuffle
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
