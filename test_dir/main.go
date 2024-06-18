package main

import "fmt"

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
}
