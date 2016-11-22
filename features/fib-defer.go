package main

import "fmt"

func main() {
	fibonacci := NewFibGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println("Fib(", i, ") =", fibonacci())
	}
}

func NewFibGenerator() func() int {
	a, b := 0, 1
	return func() int {
		defer func() { a, b = b, a+b }() // HL
		return a
	}
}

// OMIT
