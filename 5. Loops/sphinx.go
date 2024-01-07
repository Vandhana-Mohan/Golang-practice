package main

import (
  "fmt"
)

func ask() (int) {
  var input int
  fmt.Print("I am thinking of a number between 1-100: ")
  fmt.Scan(&input) // Get the input from the user
  return input
}

func main() {
  var guess int 
  for guess != 56{
    guess = ask()
    if guess == 56 {
      fmt.Println("You are correct! You may go through to the treasure!")
    }
  }
}
