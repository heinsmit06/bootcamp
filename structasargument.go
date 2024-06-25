package bootcamp

import (
	"bootcamp/firststruct"

	"github.com/alem-platform/ap"
)

func PrintUserInfo(u firststruct.User) {
	Print("Name: ")
	Println(u.Name)
	if u.GetPassword() == "" {
		Println("HasPassword: no")
	} else {
		Println("HasPassword: yes")
	}
}

func Println(s string) {
	for _, r := range s {
		ap.PutRune(rune(r))
	}
	ap.PutRune('\n')
}

func Print(s string) {
	for _, r := range s {
		ap.PutRune(rune(r))
	}
}
