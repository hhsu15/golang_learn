package main

import "fmt"

// create a new type of deck, this is like creating a class
// which is a slice of string
type deck []string


// this is like adding a method to a class
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}


func (d deck) toString() string {
	[]string(d)
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	fmt.Println(cards)
	return cards
}


// return two values from a function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

