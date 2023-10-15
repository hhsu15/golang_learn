# golang

Basics for Go. Go is a procedural language - No OOP.

## Installation

https://go.dev/dl/

## Execute code

```bash
go run main.go

```

### go commands

- go run file.go
- go build file.go # this will build an executable which you can just run
- go fmt file.go # to format the code
- go install # compile and install packages
- go get # download raw sources code
- go test # run test

## variables

Decaler a variable

```go
var card string = "spade"

```

You can use the shortened version. This is telling go to figure out the type for you.

```go
card := "spade"

```

## Data types

### Array

Array is fixed length

### Slice

Slice can shrink and expand

To declare a slice:

```go
var mySlice []string = {"firstItem"}

//or
mySlice := []string {"firstItem", "secondItem"}

```

To append item to a slice

```go

append(mySlice, "secondItem")
```

## For loop

```go
for index, item := range items {
    fmt.Printlin(item)
}

```

`range` is the key we need to use whenever we want to iterate elements in a slice.

## Custom type

This is similar to use class which we create a custom type that extends from a type and we can add custom functions to it.

```go
// create a type that extends from a slice of string

// deck.go
type deck []string

```

In order to compile code from different files we neeed to run like this

```bash
go run main.go deck.go
```

To add functions to custom type,

```go
func (d deck) print() {
    for i, card := range d {
        fmt.Printli(i, card)
    }
}
```

The (d deck) is referred to as `receiver`. This basically means that:

The variable `d` is a type of `deck`. And this function `print` is attached to type `deck`. By convention it uses single letter as abreviation for variable name.

## Type conversion

To convert string into byte:

```go
greeting := "Hi there"
fmt.Println([]byte(greeting))

```

## Read and Write File to disk

Write file like this using io/ioutil

```go
import (
	"io/ioutil"
)

////
ioutil.WriteFile(filename, []byte(yourString), 0666)
```

Read file. This is normally how you would do when reading a file.

```go

bs, err := ioutil.ReadFile(filename)
if err != nil {
    // do your thing here
}

```

### Create Unitests

In the project directory, run:

```bash
go mod init cards
```

Create a test file and name "xxx_test.go"

Then you will be able to run test using:

```bash
go test
```

## Struct

Data structure in go, like a python dictionary, or javascript object.

```go
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo  // as contactInfo contactInfo
}

func main() {

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "xxx@gmai.com",
			zipCode: 10001,
		}}
    fmt.Printl(alex.firstName)

}


```

## Pointers

Go is "pass by (copied) value" language (as opposed to referecen).

Pointer operators

```go
&variable // this means give me the address (pointer) of the variable

*pointer // this means give me the values which the pointer holds
```

So again,

- use `&` to turn value into address.
- use `*` to turn address back into value.

On the other hand you see the `*` in front of where "type" should show, it means a different thing. It means that it is a pointer, with the mentioned type for the value.

```go
func(pointerToMyStruct *myStruct) updateName () {
    ...
}

```

## Value Types and Reference Types in Go

In Go, some data types referred as Value type are the ones you need to use pointers. Data types refered as Reference type, you don't need to worry about pointers. This is a big gocha in Go.

Value Types: structs, string, int, float, bool
Reference Types: slices, maps, channels, ponters, functions

This is becuase for referece types such as slice, under the hood it creates an array and a structure that records the length of the slice, the capacity of the slice, and a referecen to the underlying array. As the result, updating the slice copy also changes the underlying array data.

### Map

Map is like python dict. This is similar to `struct`.

```go
colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf456",
		"white": "#fffff",
	}

```

#### Differeces between Struct and Map

Map:

1. All keys must be the same type
2. All values must be the same type
3. Keys are index (can be iterated over)
4. Reference type

Struct:

1. Values can be different type
2. Keys don't support index (cannot be iterated over)
3. Value type

Rule of thumb:

- Use map when you are simply creating a mapping :)
- if the usecase is you don't know what keys will be there then use map. Chances are you need to create new keys on the fly
- On the other hand you know always you are going to use defined keys then Struct is better.

## Interface

Check example code how to create an interface.

```go
type bot interface {
	// this means if there is any type in this program
	// that has a function called getGreeting
	// then the type becomes a family of "bot" that can be accetpted as type "bot"
	getGreeting() string
	// if you require specific type as argumnets
	// getGreeting(string, int)(string, int) and returns
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

```

#### Concret Type and Interface Type

Concret type are types you can create with values. Such as "map", "struct", "int", "string"

Interface type you can't create with values.

## Http

Use the "net/http" package to make htto requests.

```go

package main
import (
	"fmt"
	"net/http"
	"os"
)

func main () {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	...
}
```

### Reader and Writer Interface

The response you get from http is an interface which contains Reader and Closer interface. The Body is a Reader interface which has Read method.

```
resp.Body
```

This is quite different than what you would expect from other languages. You don't receive the actual body by calling the Read method but rather you pass the slice of bytes to it to process. So a naive way might look like below code snipet.

https://pkg.go.dev/net/http#NewRequestWithContext

```go
// create a slice of byte and assign it to max 99999 space using the built-in "make" function
...
bs := make([]byte, 99999)

resp.Body.Read(bs)
fmt.Println(string(bs))
```

In practice this we would use this to log the data to console, using Copy

```go
io.Copy(os.Stdout, resp.Body) // first arg is Writer type and second is Reader.
```

To understand this better, let's create create a custom Writer following the interface signature

```go
type logWriter struct{}

// this will automatically make logWriter a Writer type since it matches the siganature
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("This is a custom writer to mimic Copy")
	return len(bs), nil
}

```

## Concurrency in Go

### Go Routine

Use Go Routine to run code in parellel

A Go Routine is essentiallu an engine that runs code line by line. By default, a programe creates one main Go Routine.

We can launch a new Go Routine (child go routine) inside the main Go Routine, by placing the keyword 'go' in front of a function call that you want to run.

```go
...
go checkLink(link)
...
```

By default, Go uses one CPU. There is something called Go Scheduler that runs one routine until it finishes or makes a blocking call (like an HTTP request). In other words the go routines are not truly running at the same time - refering to the topic "Concurrency is not parallelism"

We could use multiple core, in such case Go Schedule will distribute the Go Routines to different CPU. This will make running multiple chunks of code truly at the same time.

### Channel

By default if we don't do anything the main Go Routine will not wait for child Go Routine to finish and it might just exit the programe before the child Go Routines finish their works.

Channel is the way (the only way) to communicate across child Go Routines and the main Go Routine. Channel is like a messaging system (i.e., with send and receive message methods) that all Go Routines can access.

Syntax for sending data to chan

```go
channel <- 5 // send data "5" to channel
```

```go
myNum <- channel // receiving data from channel

fmt.Println(<- channel) // or just print the value from channel
```

When getting data from channel, it is like making an http request, it will rerun the value when the data is available (in other words if no value is ever available will hang forever).

When using Channel, it is expected to have a equal pair for sending and consumming the data in channel - 1 sending and 1 receiving, 2 sending and 2 receiving etc.

If it is out of that blanace the program will hang forever.

#### Gotcha with Channels

You should not use the same variables for differnt go routines - due to go's pass by value nature. If the varibale value get updated the go routine will result in strange behavior (since they are using the copied value rather than the original varibale).

So instead of this...

```go
...
for l := range c {
		// go routine references the outter scope varilable
		go func () {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()

	}
...

func checkLink(l string, c chan) {
    // some code that keeps updating c
    ...
}
```

You should do this:

```go
for l := range c {
		// take the value from the source and
		go func (link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)

	}

```
