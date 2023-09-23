package main
// import "fmt"
func main() {
	//var card string = "Ace of Spades"
	// card := "Ace of Spades"  // this is same as above, when defining a variable. Go will figure it out for you
	
	// a slice of type string
	// cards := []string {"Ace of Diamonds", newCard()}
	// now this can be replaced with our custom type deck
	// cards := deck {"Ace of Diamonds", newCard()} // custom type 'deck' we create
	// cards = append(cards, "Six of Spades") // append creates a new slice

	cards := newDeck()
	

	// iterate through slice
	// cards.print()

	
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
	// cards.saveToFile("my_cards")
	newCards := newDeckFromFile("my_cards")
	newCards.print()

	//shuffle 
	newCards.shuffle()
	newCards.print()

    
}

