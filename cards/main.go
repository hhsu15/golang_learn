package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	// card := "Ace of Spades"  // this is same as above, when defining a variable. Go will figure it out for you
	
	// a slice of type string
	cards := []string {"Ace of Diamonds", newCard()} 
	cards = append(cards, "Six of Spades") // append creates a new slice
	

	// iterate through slice
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Dimonds"
}