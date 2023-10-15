package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://golang.com",
		"http://amazon.com",
		"http://facebook.com",
	}

	// create a Channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}


	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	
	// for i:=0; i < len(links); i++ {
	// 	go checkLink(<-c, c)
	// }

	// forever loop
	// for {
	// 	go checkLink(<- c, c)
	// }

	// this does the same as above
	// the loop watches c, 
	// whenever channel has some value,
	// assign the value to l and run the block of code
	for l := range c {
		// function literal (anonymous function)
		// make sure you don't share the same variable between go routines which can result in unexpected behavior
		go func (link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
		
	}
}

// second argument for checkLink is a type of channel with data type of string
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
