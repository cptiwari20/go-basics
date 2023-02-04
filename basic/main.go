package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var myToken = "asdfasdf"

func main() {
	fmt.Println("Hello World!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please tell me more about yourself")

	readerString, _ := reader.ReadString('\n')

	conStiring, err := strconv.ParseFloat(strings.TrimSpace(readerString), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Got your input thanks")
		fmt.Println(conStiring)

	}

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
	// var myName string = "Chandra Prakash Tiwari"
	// myName = "Vikas tiwari"
	// fmt.Println(myName)

	// var mynumber float32 = 525566.556
	// fmt.Println(mynumber)

	// const myAnotherName = "Adarsh"
	// // myAnotherName = "Another name"
	// fmt.Println(myAnotherName)
	// fmt.Println("My token", myToken)

}
