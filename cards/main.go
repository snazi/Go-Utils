package main

import "fmt"

func main() {
	fmt.Println("Welcome to cards!")

	// NOTE: both are acceptable to declare a variable.
	// var card string = "Ace of Spades"
	card := "Ace of Spades"   // NOTE: using ":=" notation is for initializing a variable.
	card = "Five of Diamonds" // NOTE: "=" is enough to make re assignments.
	fmt.Println(card)
}

// NOTE: notation for making a function in Go
func newCard() string { // NOTE: Go compiler requires you to declare the return type. "string" after the "()" is how to define that.
	return "Five of Diamonds"
}
