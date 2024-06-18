package main

import (
	"fmt"
)

func main() {
	var s string = "Salem, World!"
	fmt.Printf("value: %v, type: %T\n", s[1], s[1])

	fmt.Println(s[:1])
}
