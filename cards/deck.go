package main

import (
	"fmt" 
	"strings"
	"io/ioutil"	
	"os"
	"math/rand"
	"time"
)

// create a new type of deck, this is like creating a class
// which is a slice of string
type deck []string


// function receivers 
// this is like adding a method to a class
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
} 



// turn a slice into a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// write to disk
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}


//cards.shuffle
func (d deck) shuffle() {
	// generate new source for true random result
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		// generate random number within the length of deck
		newPosition := r.Intn(len(d) - 1)
		// swap the current position with a random one
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}



/////////Regular functions///////
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	//fmt.Println(cards)
	return cards
}

// read from disk and turn it into a deck type
func newDeckFromFile (filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		// get out of the program
		os.Exit(1)
	}
	// take the bytes and convert it to string
	// then convert it to slice of string using split
	s := strings.Split(string(bs), ",")
	
	// then to deck
	return deck(s)
}


// return two values from a function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}


