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

		_, err := fmt.Scanf("%d", &user_input)
		if err != nil {
			continue
		}

		if random_number == user_input {
			Print("Match, you win!\n")
			return
		} else if user_input > random_number {
			Print("Lower\n")
			continue
		} else if user_input < random_number {
			Print("Higher\n")
			continue
		}
	}
}

func Print(s string) {
	for _, v := range s {
		ap.PutRune(rune(v))
	}
}
