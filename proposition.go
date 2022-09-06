package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var movies = strings.Fields("catchmeifyoucan seven titanic")

func shuffle() {
	rand.Shuffle(len(movies), func(i, j int) {
		movies[i], movies[j] = movies[j], movies[i]
	})
	fmt.Println(movies)
}
