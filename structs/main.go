package main

import "fmt"

type Address struct {
	City    string
	Pincode int
	State   string
}

type User struct {
	Name    string
	Address Address
}

func main() {
	fmt.Println("Understand structures in go Lang")
	// Objects

	var newAddress = Address{"Jabalpur", 482001, "MP"}
	var newAddress2 = Address{State: "UP"}
	fmt.Println(newAddress, newAddress2)

	// Nested Structure
	user := User{"Chandra", newAddress}
	fmt.Println(user)

	// Anonymous Stuctures
	var school = struct {
		Name string
		Add  Address
	}{
		Name: "Gurukul", Add: newAddress2}
	fmt.Println(school)

	// fetching info from a struct
	fmt.Println("Get Name of the user ", user.Name)
	fmt.Println("Get Address of the user ", user.Address.City)

}
