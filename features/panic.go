package main

import "fmt"

func panicky() {
	panic(fmt.Sprintln("Oops!")) // panic() can take an arbitrary value as argument // HL
}

func main() {
	panicky()
	fmt.Println("Exiting main")
}
