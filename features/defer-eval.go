package main

import "fmt"

func a() {
	i := 0
	defer fmt.Println(i)
	i++
}

func main() {
	a()
}
