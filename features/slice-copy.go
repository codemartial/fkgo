package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}    // len(src) == 5
	dst := make([]int, 3)          // len(dst) == 3
	copied_count := copy(dst, src) // copied_count == min(len(src), len(dst))

	fmt.Println(copied_count, dst)
}
