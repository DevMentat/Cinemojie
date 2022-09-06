package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

// The Project is a quiz for guessing movie with emojis, you have three guesses. Maybe a point system. So I need a loop.
//TODO Create a list of proposition

const numberOfGuesses int = 3

var gameName = "Cinemojie"
var answer = "catchmeifyoucan"

func main() {
	for {
		greetUsers()
		//TODO Create a function for the first proposition
		fmt.Println("First proposition:")
		propositionOne()
		userGuess := getUserInput()

		if userGuess == answer {
			fmt.Println("Congratulation! You win 3pts")
		} else if userGuess != answer {
			fmt.Println("nope, try again.")
			propositionTwo()
			userGuess2 := getUserInput2()
			try2(userGuess2)
			//NEW PROPOSITION CHANGE THE VALUE OF answer, propo1,2,3.
			shuffle()
		}
	}
}

/////// FUNKY ZONE //////////

func greetUsers() {
	fmt.Printf("Welcome to the %v game!\n The rule is simple, you have %v guesses for each movie.\n", gameName, numberOfGuesses)
}

func getUserInput() string {
	var userGuess string
	fmt.Println("Enter your guess:")
	fmt.Scan(&userGuess)
	return userGuess
}

func getUserInput2() string {
	var userGuess2 string
	fmt.Println("Enter your second guess:")
	fmt.Scan(&userGuess2)
	return userGuess2
}

func getUserInput3() string {
	var userGuess3 string
	fmt.Println("Enter your last guess:")
	fmt.Scan(&userGuess3)
	return userGuess3
}

func try2(userGuess2 string) {
	if userGuess2 == answer {
		fmt.Println("Nice, you win 2pts")
	} else if userGuess2 != answer {
		fmt.Println("Nope again... last try")
		propositionThree()
		userGuess3 := getUserInput3()
		try3(userGuess3)
	}
}

func try3(userGuess3 string) {
	if userGuess3 == answer {
		fmt.Println("Finally you get it! You win 1pt")

	} else if userGuess3 != answer {
		fmt.Printf("No.. that's bad, the answer was %v!\n", answer)
	}
}

func propositionOne() {
	emoji.Println(":airplane:")
}

func propositionTwo() {
	emoji.Println(":airplane: :man_running:")
}

func propositionThree() {
	emoji.Println(":airplane: :man_running::money_with_wings:")
}
