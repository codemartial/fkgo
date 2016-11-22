// WARNING: DO NOT GOFMT
package main

import "fmt"
import "math/rand"

var input = rand.Perm(10) // random permutation of numbers [0, 10)

func quicksort(in []int) {
	if len(in) < 2 {
		return
	}
	start, end, pivot := 0, len(in)-1, in[len(in)/2]
	for start < end {
		for ; in[start] < pivot; start++ {}
		for ; in[end] > pivot; end-- {}

		if in[start] > in[end] {
			in[start], in[end] = in[end], in[start]
		}
	}
	quicksort(in[:start])
	quicksort(in[start:])
}

func main() {
	quicksort(input)
	fmt.Println(input)

	// Sanity Check
	quicksort(nil)
}
