package main

import "fmt"

//START OMIT
func AFunction() (int, error) {
	return 0, fmt.Errorf("An error")
}

func main() {
	i := AFunction()
	fmt.Println(i)
}

//END OMIT
