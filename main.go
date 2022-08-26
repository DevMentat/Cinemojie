package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

// The Project is a quiz for guessing movie with emojis, you have three guesses. Maybe a point system. So I need a loop.
// I need to

const numberofguesses int = 3

var gamename = "Cinemojie"
var answer = "catchmeifyoucan"

func main() {

	greetUsers()
	fmt.Println("First propostion:")
	propostionOne()
	userGuess := getUserInput()
	userGuess2 := getUserInput2()
	userGuess3 := getUserInput3()

	if userGuess == answer {
		fmt.Println("Congratulation! You win 3pts")
	} else if userGuess != answer {
		fmt.Println("nope, try again.")
		propostionTwo()
		getUserInput2()
	}

	if userGuess2 == answer {
		fmt.Println("Nice, you win 2pts")

	} else if userGuess2 != answer {
		fmt.Println("Nope again... last try")
		propostionThree()
		getUserInput3()
	}
	if userGuess3 == answer {
		fmt.Println("Finally you get it! You win 1pt")

	} else if userGuess3 != answer {
		fmt.Printf("No.. that's bad, the answer was %v!\n", answer)

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

func propostionOne() {
	emoji.Println(":airplane:")
}

func propostionTwo() {
	emoji.Println(":airplane::man_running:")
}

func propostionThree() {
	emoji.Println(":airplane::man_running::money_with_wings:")
}
