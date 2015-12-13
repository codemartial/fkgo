package main

import "fmt"

var input []int = []int{4, 0, 2, 5, 1, 6, 8, 7, 9, 3}

func quicksort(in []int, done chan<- bool) {
	defer func() { done <- true }() // HL
	switch len(in) {
	default:
		pivot := len(in) / 2
		done_left, done_right := make(chan bool), make(chan bool) // HL
		go quicksort(in[:pivot], done_left)                       // HL
		go quicksort(in[pivot:], done_right)                      // HL
		<-done_left                                               // HL
		<-done_right                                              // HL
		for i := 0; i < len(in); i++ {
			if i < pivot && in[i] > in[pivot] || i > pivot && in[i] < in[pivot] {
				in[i], in[pivot] = in[pivot], in[i]
			}
		}
	case 0, 1:
		return
	}
}

func main() {
	done := make(chan bool, 1)
	quicksort(input, done)
	<-done
	fmt.Println(input)
}
