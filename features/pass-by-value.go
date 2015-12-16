package main

import "fmt"

type Player struct {
	id        uint64
	Stamina   int `json:"stam"`
	Health    int `json:"health"`
	accessTok string
}

func (p *Player) Replace(replacement *Player) {
	p = replacement
}

func main() {
	p1 := &Player{1, 100, 100, "gibberish"}
	p2 := &Player{2, 100, 100, "rubbishes"}

	p1.Replace(p2)
	if p1.id != p2.id {
		fmt.Println("Oops! Pass by value")
	} else {
		fmt.Println(p1, p2)
	}
}

// OMIT
