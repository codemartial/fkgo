// WARNING: DO NOT GOFMT
package main

import "fmt"
import "math/rand"

var input = rand.Perm(10)

func quicksort(in []int, done chan<- bool) {
	defer func() { done <- true }() // HL
	if len(in) < 2 {
		return
	}
	start, end, pivot := 0, len(in)-1, in[len(in)/2]
	for start < end {
		for in[start] < pivot { start++ }
		for in[end] > pivot { end-- }

		if in[start] > in[end] {
			in[start], in[end] = in[end], in[start]
		}
	}
	done_left, done_right := make(chan bool), make(chan bool) // HL
	go quicksort(in[:start], done_left)                       // HL
	go quicksort(in[start:], done_right)                      // HL
	_, _ = <-done_left, <-done_right                          // HL
}

func main() {
	done := make(chan bool, 1)
	quicksort(input, done)
	<-done
	fmt.Println(input)
	// Sanity Check
	quicksort(nil, done)
	<-done
}
