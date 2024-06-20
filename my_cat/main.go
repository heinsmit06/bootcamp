package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	filename := os.Args[1:]
	if len(filename) == 0 {
		for _, v := range "my_cat file ..." {
			ap.PutRune(v)
		}
		return
	}

	for _, v := range filename {
		content, _ := os.ReadFile(v)
		if content == nil {
			s := "my_cat: " + v + ": No such file or directory\n"
			for _, g := range s {
				ap.PutRune(g)
			}
			return
		} else {
			for _, m := range content {
				ap.PutRune(rune(m))
			}
		}
	}
}
