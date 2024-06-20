package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		for _, v := range args[i] {
			ap.PutRune(rune(v))
		}
	}
}

//func main() {
//	args := os.Args[1:] // Exclude the first argument which is the program name
//
//	for _, arg := range args {
//		// Output each argument followed by a newline
//		for _, char := range arg {
//			ap.PutRune(char)
//		}
//		ap.PutRune('\n')
//	}
//}
