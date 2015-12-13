package main

import "fmt"

func a() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	fmt.Println(a())
}
