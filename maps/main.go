package main

import "fmt"

func main () {

	// to create a map
	// var colors map[string]string
	
	// make an empty map
	// colors := make(map[string]string)
	// colors["white"] = "#afefe0"
	//create a map with key as string and value as string
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf456",
		"white": "#fffff",
	}
	// iterate over a map
    printMap(colors)
	
}

func printMap (c map[string]string) {
    // iterate over a map
	for key, val := range c {
		fmt.Println("Hex code for ", key,  "is ", val)
	}

}