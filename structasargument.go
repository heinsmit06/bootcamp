package bootcamp

import (
	"bootcamp/firststruct"
	"fmt"
)

func PrintUserInfo(u firststruct.User) {
	fmt.Println("Output:")
	fmt.Println("Name: %v", u.Name)
	if u.GetPassword() == "" {
		fmt.Println("HasPassword: no")
	} else {
		fmt.Println("HasPassword: yes")
	}
}
