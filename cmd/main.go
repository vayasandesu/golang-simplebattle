package main

import (
	"fmt"
	"simplebattle/combat"
	"simplebattle/command"
)

var characters []combat.IDamagable

func Log(logable combat.ILogable) {
	fmt.Println(logable.GetInformation())
}

func ExecuteCommand(cmd command.ICommand) {
	cmd.Execute()
	if(characters == nil) {
		return;
	}

	for _, character := characters {
		Log(character)
	}
}

func main() {
	playerA := combat.Character{
		Name: "Hero",
		Hp:   10,
		Atk:  3,
		Def:  1,
	}
	playerB := combat.Monster{
		Name: "Slime",
		Hp:   5,
		Atk:  1,
		Def:  0,
	}
	playerA.Attack(&playerB)
	Log(playerA)
	Log(playerB)
}
