package bootcamp

import "github.com/alem-platform/ap"

func PrintBooks(books []*Book) {
	var name_len int = len("Name")
	var author_len int = len("Author")

	for i := range books {
		if len(books[i].Name) > name_len {
			name_len = len(books[i].Name)
		}
		if len(books[i].Author) > author_len {
			author_len = len(books[i].Author)
		}
	}

	PrintString(AddSpaces("Name", name_len))
	PrintString(AddSpaces("Author", author_len))
	PrintStringln("Year")

	for i := range books {
		PrintString(AddSpaces(books[i].Name, name_len))
		PrintString(AddSpaces(books[i].Author, author_len))
		PrintStringln(Itoa(books[i].Year))
	}
}

func AddSpaces(s string, n int) string {
	// var res string = s
	for len(s) <= n {
		s += " "
	}
	return s
}

func PrintString(s string) {
	for _, v := range s {
		ap.PutRune(v)
	}
}

func PrintStringln(s string) {
	for _, v := range s {
		ap.PutRune(v)
	}
	ap.PutRune('\n')
}
