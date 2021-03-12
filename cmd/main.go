package main

import (
	"fmt"
	combat "simplebattle/combat"
	hero "simplebattle/combat/hero"
	"simplebattle/general"
)

var arena combat.Arena

type PowerUp interface {
	Execute()
}

type PowerUp2 struct {
	value int
}

type PowerUp3 struct {
	value int
}

func (p PowerUp2) Execute() {
	(p).value = p.value * 2
}

func (p *PowerUp3) Execute() {
	(*p).value = p.value * 3
}

func main() {
	var array []interface{}
	array = append(array, PowerUp2{
		value: 2,
	})
	array = append(array, PowerUp3{
		value: 3,
	})

	for _, value := range array {
		var power PowerUp := &value
		power.Execute()
		// if ok {
		// 	//x := (*power)
		// 	//power.Execute()
		// 	fmt.Printf("%+v, Powerup !\n", power)
		// }
		fmt.Printf("result : %+v\n", value)
	}
	fmt.Println("===========================")

}

func CombatGame() {
	playerA := hero.Tanker{
		Name: "Hero",
		Hp:   10,
		Atk:  3,
		Def:  1,
	}
	playerB := hero.Savager{
		Name: "Slime",
		Hp:   5,
		Atk:  1,
		Def:  0,
	}

	arena = combat.Arena{
		Turn: 0,
	}

	arena.Characters = append(arena.Characters, playerA)
	arena.Characters = append(arena.Characters, playerB)
	arena.Characters = append(arena.Characters, playerB)

	commandIndex := 0
	for {
		var command general.ICommand

		character := arena.Characters[arena.Turn]
		fmt.Println(arena.GetInformation())
		fmt.Printf("TURN [%d] \n%s\n", arena.Turn, character.(general.ILogable).GetInformation())
		fmt.Println("---------------------------------")
		commands := arena.GetCommand()
		commandSize := len(commands)
		fmt.Printf("Command [%d] : \n", commandSize)
		if commandSize > 0 {
			for i, command := range commands {
				fmt.Printf("[%d] %+v\n", i, command)
			}
		}

		fmt.Printf("command : 0-%d\n", len(commands))
		fmt.Scanf("%d\n", &commandIndex)
		if commandIndex >= 0 && commandIndex < commandSize {
			command = commands[commandIndex]
		}

		arena.ExecuteCommand(command)
		fmt.Println("================================================================")
	}
}
