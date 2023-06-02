package main

import (
  "fmt"
  "math/rand"
)

func randomInt(min, max int) int {
  return min + rand.Intn(max - min)
}

func main() {
  var quotes = [...]string{
    "yoooo",
    "hi :3",
    "Hello world!",
  }

  fmt.Printf("Your quote is:%d\n",  quotes[randomInt(1, len(quotes))])
}
