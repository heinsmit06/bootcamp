package main

import (
	"fmt"
	"os"
)

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

	args := os.Args
	fmt.Printf("Program name: %v\n", args[0])
	fmt.Printf("args: %s\n", args[1:])
	filename := os.Args[1]
	fmt.Printf("filename: %s\n", filename)
	content, _ := os.ReadFile(filename)
	fmt.Println(string(content))
}
