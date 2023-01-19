package main

import "fmt"

func main() {
	fmt.Println("Understanding Arrays")

	var array1 [5]int
	array1[0] = 10
	array1[4] = 500
	fmt.Println(array1)

	var array2 [5]string
	array2[0] = "MyName"
	array2[4] = "Chandra"
	fmt.Println(array2[1])

	var array3 = [6]string{"Myname"}
	fmt.Println(len(array3))

	array4 := []string{"oneData"}
	println(len(array4))
}
