package main

import "fmt"

func a() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	a()
}
