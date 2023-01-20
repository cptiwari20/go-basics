package main

import "fmt"

func main() {
	fmt.Println("*** Understanding Slices in GO ****")
	// Abstraction over Go Arrays
	// var array = [2]int{}

	// defining slices
	//var slice []int
	var slice2 = make([]int, 5, 1000)
	slice2[0] = 10
	slice2[1] = 11
	slice2[2] = 18
	slice2[3] = 19
	slice2[4] = 101
	fmt.Println("This is slice number 2 ", slice2)

	// different funtions
	fmt.Println("This is slice 2 length ", len(slice2))
	fmt.Println("This is slice 2 capacity ", cap(slice2))

	fmt.Println("Give number before index 2 ", slice2[:2])
	fmt.Println("Give number between index 1:3 ", slice2[1:3])
	fmt.Println("Give number after index 3 ", slice2[3:])

	fmt.Println("Add more value to the slice ", append(slice2, 52))

}
