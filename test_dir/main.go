package main

import (
	"fmt"
)

type Book struct {
	Author string
	Title  string
	Pages  int
	Year   int
}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insert(newData int) {
	newNode := &Node{data: newData}
	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for (*current).next != nil {
		current = current.next
	}

	(*current).next = newNode
}

func (list *LinkedList) print() {
	current := (*list).head
	for current != nil {
		fmt.Printf("%d -> ", (*current).data)
		current = (*current).next
	}

	fmt.Println("nil")
}

func main() {
	///////////////                MAPS MAPS MAPS MAPS MAPS MAPS MAPS MAPS
	// var m map[string]int

	m1 := map[string]int{"Max": 21, "Gaisaks": 20, "Sanzhar": 69}
	fmt.Println(m1)
	fmt.Println(m1["Max"])

	m1["Aruzhan"] = 44
	fmt.Println(m1)

	delete(m1, "Gaisaks")
	fmt.Println(m1)

	for key, value := range m1 {
		fmt.Printf("key: %v, ", key)
		fmt.Printf("value: %v\n", value)
	}

	val, ok := m1["Sanzhar"]
	fmt.Printf("val: %v\n", val)
	fmt.Printf("ok: %v\n", ok)

	///////////////                os.Args, Work with files

	// args := os.Args
	// fmt.Printf("Program name: %v\n", args[0])
	// fmt.Printf("args: %s\n", args[1:])
	// filename := os.Args[1]
	// fmt.Printf("filename: %s\n", filename)
	// content, _ := os.ReadFile(filename)
	// fmt.Println(string(content))

	///////////////                Structs

	myBook := Book{
		Author: "Fitzgerald",
		Title:  "The Great Gatsby",
		Year:   1968,
		Pages:  324,
	}

	newBook := myBook.isNew()
	fmt.Printf("The Great Gatsby is new: %v", newBook)

	///////////////                Linked Lists

	list := LinkedList{}
	list.print()
	list.insert(51)
	list.insert(12)
	list.insert(66)
	list.insert(69)
	list.print()
}

func (b Book) isNew() bool {
	if b.Year >= 1970 {
		return true
	} else {
		return false
	}
}
