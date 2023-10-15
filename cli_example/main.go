package main


import (
	"fmt"
	"os"
	"io"
)

func main () {
	// print out the first argument
	fmt.Println(os.Args[1]) 
	
	// open the file
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// f is File and it's also implements Reader, just like resp.Body
	io.Copy(os.Stdout, f)
}