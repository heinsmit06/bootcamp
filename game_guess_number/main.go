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
			break
		} else if user_input > random_number {
			Print("Lower\n")
			continue
		} else if user_input < random_number {
			Print("Higher\n")
			continue
		}
	}

	return
}

func Print(s string) {
	for _, v := range s {
		ap.PutRune(rune(v))
	}
}
