package main

import "fmt"

func main() {
	fmt.Println("***UNDERSTAND POINTERS IN GO***")

	var x int = 25
	// memory address
	x = 26
	fmt.Println(&x)

	var y *int // pointer variable
	y = &x

	*y = 55

	fmt.Println("The address value of y is ", y)
	fmt.Println("The actual value of y is ", *y)
	fmt.Println("The actual value of x is ", x)

}
