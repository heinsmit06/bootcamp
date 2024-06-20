package main

import (
	"fmt"
	"math/rand"

	"github.com/alem-platform/ap"
)

func main() {
	random_number := rand.Intn(101)
	for {
		var user_input int
		Print("Guess number: ")
		fmt.Scanf("%d", &user_input)

		if random_number == user_input {
			Print("Match, you win!")
		} else if user_input > random_number {
			Print("Lower")
		} else if user_input < random_number {
			Print("Higher")
		}
	}
}

func Print(s string) {
	for _, v := range s {
		ap.PutRune(rune(v))
	}
}
