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
mySlice := []string {"firstItem"}

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
