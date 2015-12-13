package main

import "fmt"

func main() {
	m := make(map[string]int) // empty string -> int mapping

	m["one"], m["two"], m["three"] = 1, 2, 3 // insertion

	m = map[string]int{"one": 1, "two": 2, "three": 3} // initialization

	key := "one"
	val := m[key] // lookup

	if val, ok := m[key]; ok { // safe lookup
		fmt.Println(key, val)
	}

	for key, val = range m { // iteration
		fmt.Println(key, val)
	}

	delete(m, key) // deletion

	fmt.Println(m)

	// OMIT
}
