package main

import (
	"fmt"
	"simplebattle/combat"
)

func Log(logable combat.ILogable) {
	fmt.Println(logable.GetInformation())
}

func main() {
	playerA := combat.Character{
		Name: "Hero",
		Hp:   10,
		Atk:  3,
		Def:  1,
	}
	playerB := combat.Character{
		Name: "Slime",
		Hp:   5,
		Atk:  1,
		Def:  0,
	}
	playerA.Attack(&playerB)
	Log(playerA)
	Log(playerB)
}
