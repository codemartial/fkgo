package main

import "fmt"

func SayHelloTo(person string) string {
	return "Hello, " + person + "."
}

func main() {
	tahir := "طاهر" // "tahir" inferred to be a string variable // HL
	fmt.Println(SayHelloTo(tahir), "Welcome to Go Programming!")
}
