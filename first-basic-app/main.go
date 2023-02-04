package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our app.")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Tell me your name?")

	userResponse, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Sorry something went wrong!", err)
	} else {
		fmt.Println("Thanks for telling your name, ", userResponse, "!")
	}

	newReader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate our application! between 1 - 5.")
	userRatingResponse, _ := newReader.ReadString('\n')

	convertedSting, err := strconv.ParseFloat(strings.TrimSpace(userRatingResponse), 64)
	if err != nil {
		fmt.Println("Sorry something went wrong! ", err)
	} else {
		fmt.Println("Thanks for rating ", userResponse, "! Your response is ", convertedSting)
	}

}
