package main

import "fmt"

type (
	A struct{}
	B struct{ A } // B embeds (not inherits from) A // HL
)

func (this *A) PrintName() {
	fmt.Println("Type", this.GetName())
}
func (this *A) GetName() string {
	return "A"
}

func (this *B) GetName() string {
	return "B"
}

func main() {
	b := B{}
	b.PrintName() // B gets PrintName() from A // HL
}

// END OMIT
