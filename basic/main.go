package main

import "fmt"
myToken := "asdfasdf"

func main() {
	fmt.Println("Hello World!")

	// CODE == COMPILER === MACHINE LANG

	// RESERVE KEYWORDS
	// Data types
	// EVERYTHING IS A TYPE IN GO
	// STRING
	// INTEGERS
	// BOOLEAN - TRUE FALSE
	// FLOAT - 0.22 1.23 222.3
	// FUNCTIONS
	// STRUCTS
	// MAP
	// INTERFACES

	// Variables
	var myName string = "Chandra Prakash Tiwari"
	myName = "Vikas tiwari"
	fmt.Println(myName)

	var mynumber float32 = 525566.556
	fmt.Println(mynumber)

	const myAnotherName = "Adarsh"
	// myAnotherName = "Another name"
	fmt.Println(myAnotherName)
	fmt.Println("My token", myToken)

}
