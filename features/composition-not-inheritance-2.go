package main

import "fmt"

type (
	A struct{}
	B struct{ A }
)

func (a *A) PrintName() {
	fmt.Println("Type", a.GetName()) // PrintName() *always* calls a.GetName() // HL
}
func (a *A) GetName() string {
	return "A"
}

func (b *B) GetName() string {
	return "B"
}

func main() {
	b := B{}
	b.PrintName() // Delegated to A.PrintName(). No implementation inheritance // HL
}

// END OMIT
