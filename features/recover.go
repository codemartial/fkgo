package main

import "fmt"

func panicky() {
	defer func() { // HL
		if r := recover(); r != nil { // HL
			fmt.Println("Panic in panicky(). Error:", r) // HL
		} // HL
	}() // HL
	panic(fmt.Sprintln("Oops!"))
}

func main() {
	panicky()
	fmt.Println("Exiting main")
}
