package main

import "fmt"

func main() {
	fmt.Println("Understanding loops in go!")

	// newData := map[int]string{
	// 	1: "One",
	// 	2: "Onesdfg",
	// 	3: "sdfg",
	// 	4: "66",
	// 	5: "One669",
	// 	6: "One96",
	// }
	// for key, value := range newData {
	// 	fmt.Println(key, value)
	// }

	// fmt.Println(newData[1])

	// count := 0

	// for count == 0 {
	// 	fmt.Println(count)
	// 	count++
	// }

	for index := 10; index > 0; index-- {
		fmt.Println("Index", index)
	}
	for index := 0; index < 10; index++ {
		fmt.Println("Index2", index)
	}
}
