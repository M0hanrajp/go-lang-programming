package main

import "fmt"

func main() {

	var card = "Ace of spades"
	// syntax
	// var -> you are declaring a variable
	// card -> name of the variable
	// string -> data type of the variable
	// = -> asign a valid data to the variable card
	// var card = "Ace of spades" [ No error reported by the compiler ]
	// card = "Ace of spades", -> undefiend card

	// another way of decalring variables
	cardOne := "Five of hearts"
	// syntax
	// := -> create a variable name on the LHS and assign the value of RHS to it

	// When assigning a new value to previous variable then use `=`
	cardOne = "Five of diamonds"

	fmt.Println(card)
	fmt.Println(cardOne)
}

// Note:
// Go is statically typed language, https://stackoverflow.com/questions/1517582/what-is-the-difference-between-statically-typed-and-dynamically-typed-languages
// The types are checked before running (static) and the type error is immediately caught
