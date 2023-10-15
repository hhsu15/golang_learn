package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// instead of this...
	// bs := make([]byte, 99999)

	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// we do this...
	io.Copy(os.Stdout, resp.Body)
	// os.Stdout implements Writer interface
	// resp.Body implements Reader interface

	// now in order to understand the Reader and Writer interface
	// let's implement our own Reader and Writer
	// lw := logWriter{}
	// io.Copy(lw, resp.Body)
}

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("This is a custom writer to mimic Copy")
	return len(bs), nil
}
