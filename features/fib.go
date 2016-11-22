package main

import "fmt"

func main() {
	fibonacci := NewFibGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println("Fib(", i, ") =", fibonacci())
	}
}

func NewFibGenerator() func() int {
	a, b := 0, 1 // Multi-variable assignment // HL
	fib := func() int {
		a, b = b, a+b // Referencing a, b from outer scope // HL
		return a
	}
	return fib
}
