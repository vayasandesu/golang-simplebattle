package combat

import (
	"fmt"
	action "simplebattle/combat/action"
	components "simplebattle/combat/components"
	general "simplebattle/general"
)

type Arena struct {
	Turn       int
	Characters []interface{}
}

func (arena *Arena) NextTurn() {
	max := len(arena.Characters)
	turn := (arena.Turn + 1) % max
	arena.Turn = turn
}

func (arena *Arena) AddPlayer(character *interface{}) {
	arena.Characters = append(arena.Characters, character)
}

func (arena *Arena) GetInformation() string {
	log := ""
	len := len(arena.Characters)
	for i, character := range arena.Characters {
		log = log + fmt.Sprintf("[%d] ", i) + character.(general.ILogable).GetInformation()
		if i+1 < len {
			log = log + "\n"
		}
	}
	return log
}

func (arena *Arena) ExecuteCommand(command general.ICommand) {
	if command != nil {
		command.Execute()
	}

	for _, character := range arena.Characters {
		character.(general.ILogable).GetInformation()
	}
	arena.NextTurn()
}

func (arena Arena) GetCommand() []general.ICommand {
	var commands []general.ICommand
	character := arena.Characters[arena.Turn]

	commands = addDefenseCommand(character, commands)
	commands = addAttackCommand(character, commands)
	commands = addRestoreCommand(character, commands)

	return commands
}

func addDefenseCommand(character interface{}, commands []general.ICommand) []general.ICommand {
	component, ok := (character).(components.IDefender)

	if ok {
		command := action.DefenseCommand{
			Self: &component,
		}
		return append(commands, command)
	}

	return commands
}

func addAttackCommand(character interface{}, commands []general.ICommand) []general.ICommand {
	component, ok := (character).(components.IAttacker)
	if ok {
		command := action.AttackCommand{
			Attacker: &component,
			Target:   nil,
		}
		return append(commands, command)
	}

	return commands
}

func addRestoreCommand(character interface{}, commands []general.ICommand) []general.ICommand {
	component, ok := (character).(components.IRestorer)

	if ok {
		command := action.RestoreCommand{
			Self: &component,
		}
		return append(commands, command)
	}

	return commands
}
