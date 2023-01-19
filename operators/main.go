package main

import "fmt"

func main() {
	fmt.Println("Get started with Operators in Go")
	// Operators
	// Go - Operators
	// + - / * ++ --
	var count int = 4
	var x int = 5
	fmt.Println(count % x)

	// == != > < <= >=
	fmt.Println(count <= x)

	// && || !
	var name string = "Amit"
	var name2 string = "Sumit"
	fmt.Println(!name && !name2)

}
