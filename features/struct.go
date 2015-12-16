package main

import "fmt"

type Player struct {
	id        uint64
	Stamina   int `json:"stam"`   // Struct tag
	Health    int `json:"health"` // Struct tag
	accessTok string
}

func main() {
	var p0 Player                          // Declaration, zero-initialised
	p1 := Player{1, 100, 100, "gibberish"} // Initialiser
	p2 := Player{id: 2}                    // Partial initialiser
	p2.accessTok = "moregibberish"         // Field assignment

	fmt.Println(p0)
	fmt.Println(p1)
	fmt.Println(p2)
}

// OMIT
