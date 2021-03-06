package main

import "fmt"

type Player struct {
	id        uint64
	Stamina   int `json:"stam"`
	Health    int `json:"health"`
	accessTok string
}

func (attacker *Player) Attack(defender *Player) error { // HL
	if attacker.Stamina < 5 {
		return fmt.Errorf("Attacker %d is out of stamina", attacker.id)
	}
	if defender.Health < 5 {
		return fmt.Errorf("Defender %d is iced", defender.id)
	}
	attacker.Stamina, defender.Health = attacker.Stamina-1, defender.Health-1
	return nil
}

func main() {
	p1 := &Player{1, 100, 100, "gibberish"} // HL
	p2 := &Player{2, 100, 100, "rubbishes"} // HL

	p1.Attack(p2)
	if p1.Stamina == 100 && p2.Health == 100 {
		fmt.Println("Oops! Pass by value")
	} else {
		fmt.Println(p1, p2)
	}
}

// OMIT
