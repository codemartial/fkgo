package main

import "fmt"

func Fibbernachi() func() int {
	a, b := 0, 1 // Multi-variable assignment // HL
	fib := func() int {
		a, b = b, a+b // Referencing a, b from outer scope // HL
		return a
	}
	return fib
}

func main() {
	fibonacci := Fibbernachi()
	for i := 0; i < 10; i++ {
		fmt.Println("Fib(", i, ") =", fibonacci())
	}
}
