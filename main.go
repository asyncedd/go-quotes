package main

import (
	"fmt"
	"math/rand"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var numberOfQuotes int

	quotes := [...]string{
		"yoooo",
		"hi :3",
		"Hello world!",
	}

	fmt.Printf("Enter the amount of quotes you want: ")
	_, err := fmt.Scanln(&numberOfQuotes)
	if err != nil {
		fmt.Printf("Invalid input. Please enter a valid integer")
		return
	}

	if numberOfQuotes < 0 {
		fmt.Printf("Invalid number. Please enter a positive integer")
		return
	}

	for i := 0; i < numberOfQuotes; i++ {
		fmt.Printf("Your quote is: %s\n", quotes[randomInt(0, len(quotes))])
	}
}
