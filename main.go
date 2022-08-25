package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

// The Project is a quiz for guessing movie with emojis, you have three guesses. Maybe a point system. So I need a loop.
// I need to

const numberofguesses int = 3

var gamename = "Cinemojie"

func main() {

	greetUsers()
	fmt.Println("First propostion:")
	emoji.Println(":monkey::earth_americas:")
	for {
		userGuess := getUserInput()
		answer := "planetoftheapes"
		if userGuess == answer {
			fmt.Println("Congratulation!!")
			break
		} else {
			fmt.Println("nope, try again.")
		}

	}
}

/////// FUNKY ZONE //////////

func greetUsers() {
	fmt.Printf("Welcome to the %v game!\n The rule is simple, you have %v guesses for each movie.\n", gamename, numberofguesses)
}

func getUserInput() string {
	var userGuess string
	fmt.Println("Enter your guess:")
	fmt.Scan(&userGuess)
	return userGuess
}
