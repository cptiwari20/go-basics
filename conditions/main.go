package main

import "fmt"

func main() {
	fmt.Println("Understand Conditional Statements in Go")

	// Agar mera name "Chandra"
	// print CHandra

	// Tum Chandra nahi ho

	// var name string = "Aditya"

	if !false {
		fmt.Println("Hey Chandra, Welcome to your course! ")
	} else if false {
		fmt.Println("Hey Sumit, Welcome to your course! ")
	} else {
		fmt.Println("Hey You are not Chandra and Sumit, still welcome to the course")
	}

}
