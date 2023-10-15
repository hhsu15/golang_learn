package main

import "fmt"

type bot interface {
	// this means if there is any type in this program
	// that has a function called getGreeting
	// then the type becomes a family of "bot" that can be accetpted as type "bot"
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// very custom logic
	return "hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
