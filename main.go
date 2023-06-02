package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type Quote struct {
	Text   string `json:"q"`
	Author string `json:"a"`
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var numberOfQuotes int

	quotes := []string{
		"yoooo \n - Asyncedd",
		"hi :3 \n - Also me",
		"Hello world! \n - Some dude",
	}

	fmt.Printf("Enter the amount of quotes you want: ")
	_, err := fmt.Scanln(&numberOfQuotes)
	if err != nil {
		fmt.Printf("Invalid input. Please enter a valid integer")
		return
	}

	resp, err := http.Get("https://zenquotes.io/api/quotes")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	var quotesAPI []Quote
	err = json.NewDecoder(resp.Body).Decode(&quotesAPI)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	for _, q := range quotesAPI {
		quote := fmt.Sprintf("%s \n - %s", q.Text, q.Author)
		quotes = append(quotes, quote)
	}

	if numberOfQuotes < 0 {
		fmt.Printf("Invalid number. Please enter a positive integer")
		return
	}

	for i := 0; i < numberOfQuotes; i++ {
		fmt.Printf("Your quote is: %s\n", quotes[randomInt(0, len(quotes))])
	}
}
