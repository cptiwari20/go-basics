package main

import (
	"fmt"
)

type User struct {
	Name    string
	Age     int
	Address string
}

type count int

func main() {
	fmt.Println("Welcome to GO - next generation Programming language")

	newUser := User{"Vikas", 25, "India"}
	user2 := User{"user2", 18, "US"}
	user3 := User{"user3", 26, "China"}
	user4 := User{"user4", 55, "Russia"}

	newUser.greet()
	user2.greet()
	user3.greet()
	user4.greet()

	counter := count(20)
	fmt.Println(counter.add(56))

}

func (user User) greet() {
	fmt.Println("Hello there", user.Name)
}

func (c count) add(num count) count {
	return c + num
}
