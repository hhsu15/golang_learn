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

## For loop

```go
for index, item := range items {
    fmt.Printlin(item)
}

```

`range` is the key we need to use whenever we want to iterate elements in a slice.
