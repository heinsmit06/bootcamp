package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	args := os.Args
	filename := args[0]

	var n int

	for i := len(filename) - 1; i >= 0; i-- {
		if rune(filename[i]) == '/' || rune(filename[i]) == '.' {
			n = i + 1
		}
	}

	file_name := filename[n:]

	for i := 0; i < len(file_name); i++ {
		ap.PutRune(rune(file_name[i]))
	}
}

///////////////// NOT MINE

//func main() {
//	args := os.Args[0]
//
//	var n int
//	for i := len(args) - 1; i >= 0; i-- {
//		if args[i] == '.' || args[i] == '/' {
//			n = i + 1
//			break
//		}
//	}
//	str := args[n:]
//
//	for _, char := range str {
//		ap.PutRune(char)
//	}
//	ap.PutRune('\n')
//}
