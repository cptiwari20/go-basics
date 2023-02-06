package main

import "fmt"

func main() {
	greet("Chandra")

	result, msg := add(5, 10)
	// myfunction()
	fmt.Println(result)
	fmt.Println(msg)

	// fmt.Println("Functions in go")
	// function execution
}

// function declaration
func myfunction() {
	fmt.Println("THis is a funtion")
}

func greet(name string) {
	fmt.Println("Hello there, how are you ", name, "?")
}

func add(num1 int, num2 int) (int, string) {
	return num1 + num2, "Thanks "
}
