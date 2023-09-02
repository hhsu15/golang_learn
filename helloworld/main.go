// package is like a project, workspace or namesapce. 
// every file that belongs to the same project need to have the same package name
package main // the name "main" is required for building executable

import "fmt"  // standard package from go

func main() {
	fmt.Println("Hi there!")
}