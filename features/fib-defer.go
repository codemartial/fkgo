package main

import "fmt"

func Fibbernacci() func() int {
	a, b := 0, 1
	return func() int {
		defer func() { a, b = b, a+b }() // HL
		return a
	}
}

func main() {
	fibonacci := Fibbernacci()
	for i := 0; i < 10; i++ {
		fmt.Println("Fib(", i, ") =", fibonacci())
	}
}

// OMIT
