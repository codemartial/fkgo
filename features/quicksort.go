package main

import "fmt"

var input []int = []int{4, 0, 2, 5, 1, 6, 8, 7, 9, 3}

func quicksort(in []int) {
	switch len(in) {
	default:
		pivot := len(in) / 2
		quicksort(in[:pivot])
		quicksort(in[pivot:])
		for i := range in {
			if i < pivot && in[i] > in[pivot] || i > pivot && in[i] < in[pivot] {
				in[i], in[pivot] = in[pivot], in[i]
			}
		}
	case 0, 1: // Look Ma, No Fallthrough!
		return
	}
}

func main() {
	quicksort(input)
	fmt.Println(input)
}
