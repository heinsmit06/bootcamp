package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	args := os.Args
	num_args := len(args) - 1
	num_args_str := ""

	if num_args == 0 {
		num_args_str = "0"
	} else {

		for num_args > 0 {
			digit := byte(num_args % 10)
			num_args_str = string(digit) + num_args_str
			num_args /= 10
		}

		for i := 0; i < len(num_args_str); i++ {
			ap.PutRune(rune(num_args_str[i]))
		}

	}

	ap.PutRune('\n')
}

//func main() {
//	args := os.Args
//
//	numArgs := len(args) - 1
//
//	var numArgsStr string
//	if numArgs == 0 {
//		numArgsStr = "0"
//	} else {
//		numArgsStr = ""
//		for numArgs > 0 {
//
//			digit := byte('0' + numArgs%10)
//
//			numArgsStr = string(digit) + numArgsStr
//
//			numArgs /= 10
//		}
//	}
//
//	for _, char := range numArgsStr {
//		ap.PutRune(char)
//	}
//
//	ap.PutRune('\n')
//}
//
