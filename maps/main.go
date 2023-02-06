package main

import "fmt"

func main() {
	fmt.Println("Understanding Maps in GoLang")
	// hash tables - Objects

	// Map === Object
	// Struct === Advance Object = Classes

	// { key: value, key2: value2 }

	newMap := map[int]string{}

	newMap[10] = "Ten"
	newMap[9] = "Nine"
	newMap[8] = "Eight"
	newMap[7] = "Seven"

	fmt.Println(newMap[7])

	anotherMap := make(map[string]int)
	anotherMap["Name"] = 456
	anotherMap["age"] = 25

	delete(anotherMap, "Name")

	fmt.Println("AnotherMAP ==", anotherMap["Name"])

}
