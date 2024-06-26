package main

import (
	"bufio"
	"fmt"
	"os"
)

type Book struct {
	Name   string
	Author string
	Year   int
}

func GetBooksFromCsv(path string) []*Book {
	books := []*Book{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstLine := true
	csv_format := []string{}

	for scanner.Scan() { // scans every new line until the end
		line := scanner.Text()

		if firstLine {
			csv_format = parseLine(line)
			firstLine = false
			continue
		}

		splittedFile := parseLine(line)
		book := &Book{}

		for i, col := range csv_format {
			switch col {
			case "Name":
				book.Name = splittedFile[i]
			case "Author":
				book.Author = splittedFile[i]
			case "Year":
				book.Year = Atoi6(splittedFile[i])
			}
		}

		books = append(books, book)
	}

	return books
}

func main() {
	books := GetBooksFromCsv("books.csv")
	for _, b := range books {
		fmt.Println(b.Name, b.Author, b.Year)
	}

	//books2 := GetBooksFromCsv("books2.csv")
	//for _, b := range books2 {
	//	fmt.Println(b.Name, b.Author, b.Year)
	//}
}

func Atoi6(s string) int {
	var res int
	power := 0
	for i := len(s) - 1; i >= 0; i-- {
		res += int(s[i]-48) * PowRecursion6(10, power)
		power++
	}

	return res
}

func PowRecursion6(x int, power int) int {
	if power < 0 {
		return -1
	} else if power == 0 {
		return 1
	}

	var result int = 1
	if power > 0 {
		result = x * PowRecursion6(x, power-1)
	}
	return result
}

func parseLine(line string) []string {
	splittedFile := []string{}
	start := 0

	for i := range line {
		if line[i] == ',' {
			splittedFile = append(splittedFile, line[start:i])
			start = i + 1
		}

		splittedFile = append(splittedFile, line[start:])
	}

	return splittedFile
}
