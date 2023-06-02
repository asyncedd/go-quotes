package main

import (
  "fmt"
  "math/rand"
)

func randomInt(min, max int) int {
  return min + rand.Intn(max - min)
}

func main() {
  var numberOfQuotes int

  var quotes = [...]string{
    "yoooo",
    "hi :3",
    "Hello world!",
  }

  fmt.Printf("Enter the amount of quotes you want: ")
  fmt.Scanln(&numberOfQuotes)

  for i := 0; i < numberOfQuotes; i++ {
    fmt.Printf("Your quote is: %s\n",  quotes[randomInt(0, len(quotes))])
  }
}
