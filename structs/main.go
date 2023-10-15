package main

import "fmt"

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
	//alex := person{"Alex", "Anderson"}

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "xxx@gmai.com",
			zipCode: 10001,
		}}
	//fmt.Println(alex.firstName)
	//fmt.Printf("%+v", alex) // use Printf to print the property names

	// or upi can do this too
	// var alex person // to declare a struct, by default with 0 value (like "", 0, false)
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
    
	// use the pointer...
	// alexPointer := &alex // get the address of alex
	// alexPointer.updateName("Alexendra")

	// but go allows you to use a shortcut like this
	alex.updateName("Alexendra")
	alex.print()
}

// you will be able to do the same as receiver

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(name string) {
	
	// here we are actually getting a copy of the person 
	(*pointerToPerson).firstName = name
	
}

